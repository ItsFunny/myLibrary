# 涉及部分AI


- FFMPEG
    -   截取部分视频:
        > [参考链接](https://blog.csdn.net/huangxingli/article/details/46663143)
        ffmpeg  -i ./plutopr.mp4 -vcodec copy -acodec copy -ss 00:00:10 -to 00:00:15 ./cutout1.mp4 -y
        - 可能的情况:
            -   -ss 起始位置不存在,结果: **截取下来的视无法播放**
            -   起始位置加上-t的参数超过视频长度, 结果: **无需担心,视频自动截取到末尾结束**
        -   参数:
            
    -   将mp4转为m3u8: `ffmpeg -i input.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8`

        -   参数:
             -   -hls_time : 指定生成的每个ts切片的时长