package com.charile.exception;

import com.charile.blockchain.ResultInfo;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 11:03
 */
public class ConfigException extends  BaseBussException
{
    public ConfigException(Exception e, ResultInfo resultInfo)
    {
        super(e, resultInfo);
    }

    public ConfigException(String msg, ResultInfo resultInfo)
    {
        super(msg, resultInfo);
    }

    public ConfigException(ResultInfo resultInfo)
    {
        super(resultInfo);
    }

    /**
     * Constructs a new runtime exception with the specified detail message.
     * The cause is not initialized, and may subsequently be initialized by a
     * call to {@link #initCause}.
     *
     * @param message the detail message. The detail message is saved for
     *                later retrieval by the {@link #getMessage()} method.
     * @param code
     * @param desc
     */
    public ConfigException(String message, Integer code, String desc)
    {
        super(message, code, desc);
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
     * @param code
     * @param desc
     * @since 1.4
     */
    public ConfigException(String message, Throwable cause, Integer code, String desc)
    {
        super(message, cause, code, desc);
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
     * @param code
     * @param desc
     * @since 1.4
     */
    public ConfigException(Throwable cause, Integer code, String desc)
    {
        super(cause, code, desc);
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
     * @param code
     * @param desc
     * @since 1.7
     */
    public ConfigException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace, Integer code, String desc)
    {
        super(message, cause, enableSuppression, writableStackTrace, code, desc);
    }

    /**
     * Constructs a new runtime exception with {@code null} as its
     * detail message.  The cause is not initialized, and may subsequently be
     * initialized by a call to {@link #initCause}.
     *
     * @param code
     */
    public ConfigException(Integer code)
    {
        super(code);
    }

    /**
     * Constructs a new runtime exception with the specified detail message.
     * The cause is not initialized, and may subsequently be initialized by a
     * call to {@link #initCause}.
     *
     * @param message the detail message. The detail message is saved for
     *                later retrieval by the {@link #getMessage()} method.
     * @param code
     */
    public ConfigException(String message, Integer code)
    {
        super(message, code);
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
     * @param code
     * @since 1.4
     */
    public ConfigException(String message, Throwable cause, Integer code)
    {
        super(message, cause, code);
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
     * @param code
     * @since 1.4
     */
    public ConfigException(Throwable cause, Integer code)
    {
        super(cause, code);
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
     * @param code
     * @since 1.7
     */
    public ConfigException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace, Integer code)
    {
        super(message, cause, enableSuppression, writableStackTrace, code);
    }


    public ConfigException(Integer code, String msg)
    {
        super(code, msg);
    }

    public ConfigException(String msg)
    {
        super(msg);
    }

    public ConfigException(Exception e)
    {
        super(e);
    }

    public ConfigException(String msg, Throwable t)
    {
        super(msg, t);
    }
}
