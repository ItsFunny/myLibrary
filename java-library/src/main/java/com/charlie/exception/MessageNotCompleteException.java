package com.charlie.exception;

/**
 * Created by lizhongwen on 2018/11/15.
 */
public class MessageNotCompleteException extends RuntimeException{

    private static final long serialVersionUID = 1L;

    public MessageNotCompleteException() {
        super();
    }

    public MessageNotCompleteException(String message, Throwable cause) {
        super(message, cause);
    }

    public MessageNotCompleteException(String message) {
        super(message);
    }

    public MessageNotCompleteException(Throwable cause) {
        super(cause);
    }
}
