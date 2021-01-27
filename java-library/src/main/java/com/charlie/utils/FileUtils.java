package com.charlie.utils;

import com.charlie.blockchain.util.StringUtils;
import com.charlie.model.TwoTuple;
import org.apache.commons.io.IOUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.activation.MimetypesFileTypeMap;
import java.io.*;
import java.nio.channels.FileChannel;
import java.util.*;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-09 15:26
 */
public class FileUtils
{
    private static Logger logger = LoggerFactory.getLogger(FileUtils.class);

    private static final int DEFAULT_BUFFER_SIZE = 1024 * 1024;

    public static final int ENTERPRISE_QR_CODE_BUFFER_SIZE = 16 * 1024 * 1024;
    // 1m大小
    private static final int _1_M = 1024 * 1024;

    private static final int _1_k = 1024;

    public static String getSuffix(String path)
    {
        return path.substring(path.lastIndexOf(".") + 1);
    }


    public static byte[] readFileToByteArray(File file) throws IOException
    {
        FileInputStream in = null;

        byte[] var2;
        try
        {
            in = openInputStream(file);
            var2 = IOUtils.toByteArray(in, file.length());
        } finally
        {
            IOUtils.closeQuietly(in);
        }

        return var2;
    }

    public static FileInputStream openInputStream(File file) throws IOException
    {
        if (file.exists())
        {
            if (file.isDirectory())
            {
                throw new IOException("File '" + file + "' exists but is a directory");
            } else if (!file.canRead())
            {
                throw new IOException("File '" + file + "' cannot be read");
            } else
            {
                return new FileInputStream(file);
            }
        } else
        {
            throw new FileNotFoundException("File '" + file + "' does not exist");
        }
    }
    // 拆分文件

    /**
     * Split files.
     *
     * @param filePath    the file path 原始的文件路径
     * @param size        the size  代表按多少m切割文件
     * @param toDirectory the to directory  存储的目录,不存在则会创建
     * @param delete      the delete    是否删除原先文件
     * @throws IOException the io exception
     */
    public static void splitFiles(String filePath, int size, String toDirectory, boolean delete) throws IOException
    {
        File originFile = new File(filePath);
        if (!originFile.exists())
        {
            throw new RuntimeException("该文件不存在,文件路径为:" + filePath);
        }

        String originName = originFile.getName();
        FileInputStream inputStream = new FileInputStream(originFile);
        FileChannel inChannel = inputStream.getChannel();
        FileOutputStream out = null;
        FileChannel outChannel = null;

        if (!toDirectory.endsWith(File.separator))
        {
            toDirectory += File.separator;
        }
        // 每个chunk的大小
        long chunkSize = size * _1_k;
        // 计算最终会分成几个文件
        long totalLength = originFile.length();
        int count = (int) (totalLength / chunkSize);
        if (createDirIfNotExist(toDirectory))
        {
            logger.debug("创建文件夹:{} 成功", toDirectory);
        }
        String chunkFileName = toDirectory + originName + "-";
        try
        {
            for (int i = 0; i <= count; i++)
            {
                // 生成文件的路径
                String newChunkName = chunkFileName + i;
//                System.out.println("创建文件:" + newChunkName);
                out = new FileOutputStream(new File(newChunkName));
                outChannel = out.getChannel();
                // 从inChannel的m*i处，读取固定长度的数据，写入outChannel
                if (i != count)
                {
                    inChannel.transferTo(chunkSize * i, chunkSize, outChannel);
//                   不可取 outChannel.transferFrom(inChannel, i * chunkSize, chunkSize);
                } else// 最后一个文件，大小不固定，所以需要重新计算长度
                {
                    inChannel.transferTo(chunkSize * i, totalLength - chunkSize * count, outChannel);
//                    outChannel.transferFrom(inChannel, chunkSize * count, totalLength - chunkSize * count);
                }
                out.close();
                outChannel.close();
            }
            if (delete)
            {
                logger.info("删除原先的文件");
                originFile.delete();
            }
        } finally
        {
            inputStream.close();
            inChannel.close();
        }
    }

    // 创建文件夹
    public static boolean createDirIfNotExist(String dir)
    {
        // 判断是否以 / 结尾
        if (!dir.endsWith(File.separator))
        {
            dir += File.separator;
        }
        File file = new File(dir);
        if (file.exists())
        {
            logger.debug("文件夹:" + dir + " 已经存在");
            return true;
        }
        return file.mkdirs();
    }

    // 合并文件:
    public static void mergeFile(String directoryPath, String outputFile) throws IOException
    {
        File directory = new File(directoryPath);
        if (!directory.isDirectory()) throw new RuntimeException("该路径:" + directoryPath + " 非文件夹路径,必须为文件夹路径");
        File[] files = directory.listFiles();
        List<File> fileList = Arrays.asList(files);
        Collections.sort(fileList, new Comparator<File>()
        {
            @Override
            public int compare(File o1, File o2)
            {
                String name = o1.getName();
                String[] split = name.split("-");
                int index1 = Integer.parseInt(split[split.length - 1]);
                name = o2.getName();
                split = name.split("-");
                int index2 = Integer.parseInt(split[split.length - 1]);
                return index1 <= index2 ? -1 : 1;
            }
        });
        FileOutputStream fileOutputStream = null;
        FileChannel outChannel = null;
        FileChannel inChannel = null;
        FileInputStream inputStream = null;

        try
        {
            fileOutputStream = new FileOutputStream(new File(outputFile));
            outChannel = fileOutputStream.getChannel();

            long start = 0l;
            for (File file : fileList)
            {
                inputStream = new FileInputStream(file);
                inChannel = inputStream.getChannel();
                outChannel.transferFrom(inChannel, start, file.length());
                logger.debug("合并文件:" + file.getName());
                start += file.length();

                inputStream.close();
                inChannel.close();
                if (!file.delete())
                {
                    logger.error("删除文件:{}失败", file.getName());
                }
            }
        } catch (IOException e)
        {
            e.printStackTrace();
            logger.error("文件发送io错误:" + e.getMessage());
            throw e;
        } finally
        {
            try
            {
                if (null != fileOutputStream) fileOutputStream.close();
                if (null != outChannel) outChannel.close();
            } catch (IOException e)
            {
                e.printStackTrace();
            }
        }
    }


    public static String appendFilePathIfNone(String path)
    {
        if (!path.endsWith(File.separator))
        {
            path += File.separator;
        }
        return path;
    }

    public static String cutPathIfStartWith(String path)
    {
        if (path.startsWith(File.separator))
        {
            path = path.substring(1);
        }
        return path;
    }




    /**
     * 校验文件
     *
     * @param path 文件路径
     * @param flag 当文件不存在时是否创建文件 [true: 创建文件；false: 抛出文件不存在异常]
     * @return
     * @throws Exception
     */
    public static File checkFilePath(String path, boolean flag) throws Exception {
        if (StringUtils.isBlank(path)) {
            throw new RuntimeException("The file path cannot be empty.");
        }
        File file = new File(path);
        if (!file.exists()) {
            if (flag) {
                // ----- 当文件不存在时，创建新文件
                if (!file.createNewFile()) {
                    throw new RuntimeException("Failed to create file.");
                }
            } else {
                // ----- 抛出文件不存在异常
                throw new RuntimeException("File does not exist.");
            }
        }
        return file;
    }

    public static String getFilePath(String file) {
        return getFilePath(new File(file));
    }

    public static String getFilePath(File file) {
//		if (!file.exists()) {
//			return null;
//		}
        String absolutePath = file.getAbsolutePath();
        return absolutePath.substring(0, absolutePath.length() - (file.getName().length() + File.separator.length()));
    }

    public static String getFilename(String file) {
        return getFilename(new File(file));
    }

    public static String getFilename(File file) {
//		if (!file.exists()) {
//			return null;
//		}
        return file.getName();
    }

    public static boolean checkExist(String fileName) {
        if (StringUtils.isNullEmpty(fileName)) {
            return false;
        }

        File file = new File(fileName);
        return file.exists();
    }

    /**
     * 删除目录
     *
     * @author
     * @Date 2017/10/30 下午4:15
     */
    public static boolean deleteDir(File dir) {
        return deleteDir(dir, false);
    }

    /**
     * 删除目录内文件，保留目录
     *
     * @author
     * @Date 2017/10/30 下午4:15
     */
    public static boolean deleteDir(File dir, boolean keepRootDir) {
        if (dir.isDirectory()) {
            String[] children = dir.list();
            for (int i = 0; i < children.length; i++) {
                boolean success = deleteDir(new File(dir, children[i]));
                if (!success) {
                    return false;
                }
            }
        }
        if (!keepRootDir) {
            return dir.delete();
        }
        return true;
    }

    public static String getAbsolutePath(String path) {
        try {
            File file = new File(path);
            if (!file.exists()) {
                return null;
            }
            return file.getCanonicalPath();
        } catch (IOException e) {
            return null;
        }
    }

    /**
     * 返回一个相对于relatedPath的路径 如果path为绝对路径，那么使用path
     * 如果path为相对路径，那么返回相对于releatedPath的绝对路径 *
     */
    public static String getAbsolutePathRelated(String relatedPath, String path) {
        if (new File(path).isAbsolute())
            return path;

        File relatedFile = new File(relatedPath);
        if (relatedFile.isDirectory()) {
            try {
                return new File(relatedFile + "/" + path).getCanonicalPath();
            } catch (IOException e) {
                return null;
            }
        } else {
            try {
                return new File(relatedFile.getParentFile() + "/" + path).getCanonicalPath();
            } catch (IOException e) {
                return null;
            }
        }
    }

    public static Properties loadProperties(String path) {
        Properties prop = null;
        InputStream inputStream = null;
        try {
            inputStream = new FileInputStream(new File(path));
            prop = new Properties();
            prop.load(inputStream);
        } catch (IOException e) {
            IOException ex = new IOException("加载" + path + "时，文件未找到", e);
            ex.printStackTrace();
            System.exit(0);
        } finally {
            if (inputStream != null) {
                try {
                    inputStream.close();
                } catch (IOException e) {
//                    DebugUtil.println("关闭properties输入流失败，文件：" + path);
                }
            }
        }
        return prop;
    }

    /**
     *
     * 描述:这个方法是获取一个文件的文件名，不带格式
     *
     * @param
     * @return
     * @time 2016年7月22日-下午7:32:37
     */
    public static String getFileName(File file) {
        String fileName = file.getName();
        String[] names = fileName.split("\\.");
        return names[0];
    }

    /**
     *
     * 描述:将文件的内容作为String 读出
     *
     * @param pathFile
     * @return
     * @throws FileNotFoundException
     * @throws UnsupportedEncodingException
     * @time 2016年8月17日-下午3:40:41
     */
    public static String getFileText(String pathFile) throws FileNotFoundException, UnsupportedEncodingException {

        InputStream is = new FileInputStream(pathFile);
        String fileContext = "";
        BufferedReader reader = new BufferedReader(new InputStreamReader(is, "utf-8"));
        String line;
        try {
            line = reader.readLine();
            while (line != null) { // 如果 line 为空说明读完了
                fileContext += line;
                fileContext += '\n';
                line = reader.readLine(); // 读取下一行
            }
            reader.close();
            is.close();
        } catch (IOException e) {
            return "";
        } // 读取第一行

        return fileContext;
    }

    /**
     *
     * 描述:写入text到文件。如果文件存在，之前的内容将被替换。
     *
     * @param pathFile
     * @param text
     * @throws UnsupportedEncodingException
     * @throws FileNotFoundException
     * @time 2016年8月17日-下午3:42:07
     */
    public static void setFileText(String pathFile, String text)
            throws FileNotFoundException, UnsupportedEncodingException {
        PrintWriter writer = new PrintWriter(pathFile, "UTF-8");
        writer.println(text);
        writer.close();
    }

//	/**
//	 *
//	 * 描述：向一个文件中写入字节，它每次写入的时候都会覆盖原来的数据，生成一个新的文件。
//	 *
//	 * @param fileBytes
//	 *            要写入的内容
//	 * @param fileName
//	 *            要写入的文件名，包括扩展名
//	 * @param filePath
//	 *            要写入的文件路径，如果这个路径不存在，会新建一个。
//	 *
//	 *            2017年2月18日 下午5:38:35
//	 */
//	public static void createFile(byte[] fileBytes, String fileName, String filePath) {
//
//		File file = null;
//		try {
//			File dir = new File(filePath);
//			if (!dir.exists()) {// 判断文件目录是否存在
//				dir.mkdirs();
//			}
//			file = new File(filePath + "/" + fileName);
//
//			ByteArrayInputStream bin = new ByteArrayInputStream(fileBytes);
//			org.apache.commons.io.FileUtils.copyInputStreamToFile(bin, file);
//		} catch (Exception e) {
//			e.printStackTrace();
//		}
//	}

    public static byte[] fileToBytes(String path) {
        return fileToBytes(path, DEFAULT_BUFFER_SIZE);
    }

    /**
     *
     * 描述：将一个文件转化为byte数组
     *
     * @param path
     * @return
     *
     *         2017年2月20日 上午10:19:15
     */
    public static byte[] fileToBytes(String path, int bufferSize) {
        long begin = System.currentTimeMillis();
        byte[] data = null;
        File file = new File(path);
        if (!file.exists()) {
            return null;
        }

        FileInputStream fin = null;
        ByteArrayOutputStream bos = null;
        try {
            fin = new FileInputStream(file);
            bos = new ByteArrayOutputStream();
            byte[] cache = new byte[bufferSize];
            int len = -1;
            while ((len = fin.read(cache)) != -1) {
                bos.write(cache, 0, len);
            }
            data = bos.toByteArray();
        } catch (IOException e) {
            throw new RuntimeException(e);
        } finally {
            if (fin != null) {
                try {
                    fin.close();
                } catch (IOException e) {
                    throw new RuntimeException(e);
//                    LOG.warning(BSModule.COMMON, e, "关闭文件流失败， path = %s", path);
                }
            }
            if (bos != null) {
                try {
                    bos.close();
                } catch (IOException e) {
//                    LOG.warning(BSModule.COMMON, e, "关闭文件流失败， path = %s", path);
                }
            }

//            LOG.info(BSModule.COMMON, "读取文件， path = %s, 文件大小： %s, 耗时： %d", path, data == null ? 0 : StringUtil.byteToUnit(data.length),
//                    (System.currentTimeMillis() - begin));
        }
        return data;
    }

    /**
     * 根据路径创建目录
     *
     * @param pathname
     */
    public static void createPath(String pathname) {
        if (StringUtils.isNullEmpty(pathname)) {
            return;
        }
        File path = new File(pathname);
//        LOG.info(BSModule.COMMON, "create dir [%s]", path);
        if (!path.exists()) {
            path.mkdirs();
        }
    }

    public static void createFile(String filename) throws IOException {
        if (StringUtils.isNullEmpty(filename)) {
            return;
        }
        String path = filename.substring(0, filename.lastIndexOf(File.separator));
        createPath(path);
        File file = new File(filename);
        if (!file.exists()) {
            file.createNewFile();
        }
    }

    public static String joinFilePath(String... path) {
        return joinFilePathSep(File.separator, path);
    }

    public static String joinFilePathSep(String separator, String... path) {
        StringBuilder sb = new StringBuilder();
        for (String string : path) {
            sb.append(string).append(separator);
        }
        return sb.substring(0, sb.length() - separator.length());
    }

    public static byte[] readRaw(File file) throws IOException {
        FileInputStream stream = null;
        try {
            stream = new FileInputStream(file);
            byte[] data = new byte[(int) file.length()];
            stream.read(data);
            return data;
        } catch (IOException e) {
            throw e;
        } finally {
            if (stream != null) {
                try {
                    stream.close();
                } catch (IOException e1) {
                }
            }
        }
    }

    public static void writeRaw(File file, byte[] data) throws IOException {
        FileOutputStream stream = null;
        try {
            stream = new FileOutputStream(file);
            stream.write(data);
        } catch (IOException e) {
            throw e;
        } finally {
            if (stream != null) {
                try {
                    stream.close();
                } catch (IOException e1) {
                }
            }
        }
    }

    public static String readClassPathFileData(String fileName) throws IOException {
        InputStream in = FileUtils.class.getClassLoader().getResource(fileName).openStream();
        BufferedReader br = new BufferedReader(new InputStreamReader(in, "UTF-8"));
        StringBuilder sb = new StringBuilder();
        String line = "";
        while ((line = br.readLine()) != null) {
            sb.append(new String(line.getBytes(), "UTF-8") + "\n");
        }
        return sb.toString();
    }

    public static byte[] readToBytes(String fileName) throws IOException {
        File file = new File(fileName);
        Long filelength = file.length();
        byte[] filecontent = new byte[filelength.intValue()];

        try {
            FileInputStream in = new FileInputStream(file);
            in.read(filecontent);
            in.close();
        } catch (FileNotFoundException e) {
            throw new IOException(e);
        } catch (IOException e) {
            throw e;
        }

        return filecontent;
    }

    /**
     * 从文件读取字符串
     *
     * @param fileName
     * @return
     */
    public static String readToString(String fileName) throws IOException {
        byte[] filecontent = readToBytes(fileName);

        try {
            return new String(filecontent, "UTF-8");
        } catch (UnsupportedEncodingException e) {
            throw new IOException(e);
        }
    }

    /**
     * 将字符串写入文件中
     *
     * @param toFile
     * @param text
     */
    public static void writeString2File(File toFile, String text) throws IOException {
        // 先删除文件，防止数据错误
        delFile(toFile);

        // 写入数据
        OutputStreamWriter osw = null;
        try {
            if (!toFile.exists()) {
                File parent = toFile.getParentFile();
                if (parent != null && !parent.exists()) {
                    parent.mkdirs();
                }
                toFile.createNewFile();
            }

            osw = new OutputStreamWriter(new FileOutputStream(toFile), "UTF-8");
            osw.write(text);
        } catch (IOException e) {
            throw e;
        } finally {
            if (osw != null) {
                try {
                    osw.flush();
                    osw.close();
                } catch (IOException e) {
                    throw e;
                }
            }
        }
    }

    public static void writeBytes2File(String fileName, byte[] data) throws IOException {
        writeBytes2File(new File(fileName), data);
    }

    public static void writeBytes2FileAppend(String fileName, byte[] data) throws IOException {
        long begin = System.currentTimeMillis();
        try (RandomAccessFile randomFile = new RandomAccessFile(fileName, "rw");) {
            long fileLength = randomFile.length();
            randomFile.seek(fileLength);
            randomFile.write(data);
            if (randomFile.getChannel().isOpen()) {
                randomFile.getChannel().force(false);
            }
        } catch (Exception e) {
            throw new IOException(e);
        }
//        LOG.info(BSModule.COMMON, "写入文件， path = %s, 文件大小： %s, 耗时： %d", fileName, data == null ? 0 : StringUtil.byteToUnit(data.length),
//                (System.currentTimeMillis() - begin));
    }

    public static void writeBytes2File(File file, byte[] data) throws IOException {
        long begin = System.currentTimeMillis();
        // 先删除文件，防止数据错误
        delFile(file);

        // 写入数据
        try {
            if (!file.exists()) {
                File parent = file.getParentFile();
                if (parent != null && !parent.exists()) {
                    parent.mkdirs();
                }
                file.createNewFile();
            }

            try (FileOutputStream out = new FileOutputStream(file)) {
                out.write(data);
                if (out.getChannel().isOpen()) {
                    out.getChannel().force(false);
                }
            }
        } catch (IOException e) {
            throw e;
        }
//        LOG.info(BSModule.COMMON, "写入文件， path = %s, 文件大小： %s, 耗时： %d", file.getAbsolutePath(),
//                data == null ? 0 : StringUtil.byteToUnit(data.length), (System.currentTimeMillis() - begin));
    }

    public static void writeBytes2File2(String path, byte[] data) throws IOException {
        File file = new File(path);
        long begin = System.currentTimeMillis();
        // 先删除文件，防止数据错误
        delFile(file);

        // 写入数据
        FileOutputStream out = null;
        try {
            if (!file.exists()) {
                File parent = file.getParentFile();
                if (parent != null && !parent.exists()) {
                    parent.mkdirs();
                }
                file.createNewFile();
            }

            out = new FileOutputStream(file);
            out.write(data);
        } catch (IOException e) {
            throw e;
        }finally {
            if (out != null){
                out.close();
            }
        }
//        LOG.info(BSModule.COMMON, "写入文件， path = %s, 文件大小： %s, 耗时： %d", file.getAbsolutePath(),
//                data == null ? 0 : StringUtil.byteToUnit(data.length), (System.currentTimeMillis() - begin));
    }

    public static boolean delFile(String fileName) {
        return delFile(new File(fileName));
    }

    /**
     * 删除文件
     */
    public static boolean delFile(File file) {
        boolean bea = false;
        try {
            if (file.exists()) {
                file.delete();
                bea = true;
            } else {
                bea = false;
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return bea;
    }

    public static void copyFileUsingFileStreams(File source, File dest) throws IOException {
        InputStream input = null;
        OutputStream output = null;
        try {
            input = new FileInputStream(source);
            output = new FileOutputStream(dest);
            byte[] buf = new byte[1024];
            int bytesRead;
            while ((bytesRead = input.read(buf)) != -1) {
                output.write(buf, 0, bytesRead);
            }
        } finally {
            try {
                input.close();
            } catch (Exception e) {
                e.printStackTrace();
            }
            try {
                output.close();
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }

    public static String getFileMIMEType(File file) {
        return new MimetypesFileTypeMap().getContentType(file);
    }

    private static final String FILE_TYPE_PDF = "25504446";

    public static TwoTuple<Boolean, String> isPDF(byte[] dataByte) {
        final String fileHeader = getFileHeader(dataByte);
        final boolean isPdf = FILE_TYPE_PDF.equals(fileHeader);
        return new TwoTuple<>(isPdf, fileHeader);
    }

    public static String getFileHeader(byte[] dataByte) {
        if (dataByte == null || dataByte.length < 4) {
            return null;
        }
        final byte[] headerByte = Arrays.copyOfRange(dataByte, 0, 4);
        final String value = bytesToHexString(headerByte);
        return value;
    }

    /**
     * 将要读取文件头信息的文件的byte数组转换成string类型表示 下面这段代码就是用来对文件类型作验证的方法，
     * 将字节数组的前四位转换成16进制字符串，并且转换的时候，要先和0xFF做一次与运算。
     * 这是因为，整个文件流的字节数组中，有很多是负数，进行了与运算后，可以将前面的符号位都去掉，
     * 这样转换成的16进制字符串最多保留两位，如果是正数又小于10，那么转换后只有一位， 需要在前面补0，这样做的目的是方便比较，取完前四位这个循环就可以终止了
     *
     * @param src 要读取文件头信息的文件的byte数组
     * @return 文件头信息
     */
    private static String bytesToHexString(byte[] src) {
        final StringBuilder builder = new StringBuilder();
        String hv;
        for (int i = 0; i < src.length; i++) {
            // 以十六进制（基数 16）无符号整数形式返回一个整数参数的字符串表示形式，并转换为大写
            hv = Integer.toHexString(src[i] & 0xFF).toUpperCase();
            if (hv.length() < 2) {
                builder.append(0);
            }
            builder.append(hv);
        }

        final String result = builder.toString();
        return result;
    }

}
