package com.charlie.blockchain.exception;

import com.charlie.blockchain.ResultInfo;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-23 23:28
 */
public class InvokeException extends  BaseFabricException
{
    public InvokeException(Exception e, ResultInfo info)
    {
        super(e, info);
    }

    public InvokeException(ResultInfo info)
    {
        super(info);
    }

    /**
     * Constructs a new runtime exception with the specified detail message.
     * The cause is not initialized, and may subsequently be initialized by a
     * call to {@link #initCause}.
     *
     * @param message the detail message. The detail message is saved for
     *                later retrieval by the {@link #getMessage()} method.
     */
    public InvokeException(String message)
    {
        super(message);
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
    public InvokeException(String message, Throwable cause)
    {
        super(message, cause);
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
    public InvokeException(Throwable cause)
    {
        super(cause);
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
    public InvokeException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace)
    {
        super(message, cause, enableSuppression, writableStackTrace);
    }


    public InvokeException(Integer errorCode, String msg)
    {
        super(errorCode, msg);
    }
    public InvokeException(String message, Integer errorCode, String msg)
    {
        super(message, errorCode, msg);
    }

    public InvokeException(String message, Throwable cause, Integer errorCode, String msg)
    {
        super(message, cause, errorCode, msg);
    }

    public InvokeException(Throwable cause, Integer errorCode, String msg)
    {
        super(cause, errorCode, msg);
    }

    public InvokeException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace, Integer errorCode, String msg)
    {
        super(message, cause, enableSuppression, writableStackTrace, errorCode, msg);
    }
}
