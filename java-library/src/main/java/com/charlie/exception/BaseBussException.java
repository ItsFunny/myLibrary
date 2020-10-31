package com.charlie.exception;

import com.charlie.blockchain.ResultInfo;
import lombok.Data;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @Description: 基础业务异常包装类
 * @Attention:
 * @Date 创建时间：2019-12-25 13:51
 */
@Data
public class BaseBussException extends RuntimeException
{
    private Integer code;
    private String desc;

    public BaseBussException(Exception e,ResultInfo resultInfo){
        super(e);
        this.code = resultInfo.getCode();
        this.desc = resultInfo.getMsg();
    }

    public BaseBussException(String msg, ResultInfo resultInfo)
    {
        super(msg);
        this.code = resultInfo.getCode();
        this.desc = resultInfo.getMsg();
    }

    public BaseBussException(ResultInfo resultInfo)
    {
        this.code = resultInfo.getCode();
        this.desc = resultInfo.getMsg();
    }

    /**
     * Constructs a new runtime exception with the specified detail message.
     * The cause is not initialized, and may subsequently be initialized by a
     * call to {@link #initCause}.
     *
     * @param message the detail message. The detail message is saved for
     *                later retrieval by the {@link #getMessage()} method.
     */
    public BaseBussException(String message, Integer code, String desc)
    {
        super(message);
        this.code = code;
        this.desc = desc;
    }

    /**
     * Constructs a new runtime exception with the specified detail message and
     * cause.  <p>Note that the detail message associated with
     * {@code cause} is <i>not</i> automatically incorporated in
     * this runtime exception's detail message.
     *
     * @param message the detail message (which is saved for later retrieval
     *                by the {@link #getMessage()} method).
     * @param cause   the cause (which is saved for later retrieval by the
     *                {@link #getCause()} method).  (A <tt>null</tt> value is
     *                permitted, and indicates that the cause is nonexistent or
     *                unknown.)
     * @since 1.4
     */
    public BaseBussException(String message, Throwable cause, Integer code, String desc)
    {
        super(message, cause);
        this.code = code;
        this.desc = desc;
    }

    /**
     * Constructs a new runtime exception with the specified cause and a
     * detail message of <tt>(cause==null ? null : cause.toString())</tt>
     * (which typically contains the class and detail message of
     * <tt>cause</tt>).  This constructor is useful for runtime exceptions
     * that are little more than wrappers for other throwables.
     *
     * @param cause the cause (which is saved for later retrieval by the
     *              {@link #getCause()} method).  (A <tt>null</tt> value is
     *              permitted, and indicates that the cause is nonexistent or
     *              unknown.)
     * @since 1.4
     */
    public BaseBussException(Throwable cause, Integer code, String desc)
    {
        super(cause);
        this.code = code;
        this.desc = desc;
    }

    /**
     * Constructs a new runtime exception with the specified detail
     * message, cause, suppression enabled or disabled, and writable
     * stack trace enabled or disabled.
     *
     * @param message            the detail message.
     * @param cause              the cause.  (A {@code null} value is permitted,
     *                           and indicates that the cause is nonexistent or unknown.)
     * @param enableSuppression  whether or not suppression is enabled
     *                           or disabled
     * @param writableStackTrace whether or not the stack trace should
     *                           be writable
     * @since 1.7
     */
    public BaseBussException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace, Integer code, String desc)
    {
        super(message, cause, enableSuppression, writableStackTrace);
        this.code = code;
        this.desc = desc;
    }

    /**
     * Constructs a new runtime exception with {@code null} as its
     * detail message.  The cause is not initialized, and may subsequently be
     * initialized by a call to {@link #initCause}.
     */
    public BaseBussException(Integer code)
    {
        this.code = code;
    }

    /**
     * Constructs a new runtime exception with the specified detail message.
     * The cause is not initialized, and may subsequently be initialized by a
     * call to {@link #initCause}.
     *
     * @param message the detail message. The detail message is saved for
     *                later retrieval by the {@link #getMessage()} method.
     */
    public BaseBussException(String message, Integer code)
    {
        super(message);
        this.code = code;
    }

    /**
     * Constructs a new runtime exception with the specified detail message and
     * cause.  <p>Note that the detail message associated with
     * {@code cause} is <i>not</i> automatically incorporated in
     * this runtime exception's detail message.
     *
     * @param message the detail message (which is saved for later retrieval
     *                by the {@link #getMessage()} method).
     * @param cause   the cause (which is saved for later retrieval by the
     *                {@link #getCause()} method).  (A <tt>null</tt> value is
     *                permitted, and indicates that the cause is nonexistent or
     *                unknown.)
     * @since 1.4
     */
    public BaseBussException(String message, Throwable cause, Integer code)
    {
        super(message, cause);
        this.code = code;
    }

    /**
     * Constructs a new runtime exception with the specified cause and a
     * detail message of <tt>(cause==null ? null : cause.toString())</tt>
     * (which typically contains the class and detail message of
     * <tt>cause</tt>).  This constructor is useful for runtime exceptions
     * that are little more than wrappers for other throwables.
     *
     * @param cause the cause (which is saved for later retrieval by the
     *              {@link #getCause()} method).  (A <tt>null</tt> value is
     *              permitted, and indicates that the cause is nonexistent or
     *              unknown.)
     * @since 1.4
     */
    public BaseBussException(Throwable cause, Integer code)
    {
        super(cause);
        this.code = code;
    }

    /**
     * Constructs a new runtime exception with the specified detail
     * message, cause, suppression enabled or disabled, and writable
     * stack trace enabled or disabled.
     *
     * @param message            the detail message.
     * @param cause              the cause.  (A {@code null} value is permitted,
     *                           and indicates that the cause is nonexistent or unknown.)
     * @param enableSuppression  whether or not suppression is enabled
     *                           or disabled
     * @param writableStackTrace whether or not the stack trace should
     *                           be writable
     * @since 1.7
     */
    public BaseBussException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace, Integer code)
    {
        super(message, cause, enableSuppression, writableStackTrace);
        this.code = code;
    }

    public BaseBussException(Integer code, String msg)
    {
        super(msg);
        this.code = code;
    }


    public BaseBussException(String msg)
    {
        super(msg);
    }

    public BaseBussException(Exception e)
    {
        super(e);
    }

    public BaseBussException(String msg, Throwable t)
    {
        super(msg, t);
    }
}
