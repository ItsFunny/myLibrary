package cn.bidsun.blockchain.util;

import org.apache.commons.io.IOUtils;
import org.apache.commons.lang3.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.*;
import java.nio.channels.FileChannel;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

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

    public static boolean checkExist(String fileName) {
        if (StringUtils.isEmpty(fileName)) {
            return false;
        }

        File file = new File(fileName);
        return file.exists();
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

}
