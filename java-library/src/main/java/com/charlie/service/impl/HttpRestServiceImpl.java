package com.charlie.service.impl;

import com.charlie.blockchain.util.StringUtils;
import com.charlie.model.HttpRestResult;
import com.charlie.service.IHttpRestService;
import com.charlie.utils.HttpUtils;
import com.charlie.utils.JSONUtil;
import com.charlie.utils.UUIDUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.*;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.client.UnknownHttpStatusCodeException;

import java.util.Collections;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-11-22 15:59
 */
public class HttpRestServiceImpl implements IHttpRestService
{
    private Logger logger = LoggerFactory.getLogger(HttpRestServiceImpl.class);
    private static HttpRestServiceImpl instance;
//    private ThreadLocalCommandStack stack;

    @Autowired
    public void setInstance(HttpRestServiceImpl instance)
    {
        HttpRestServiceImpl.instance = instance;
    }

    public static HttpRestServiceImpl getInstance()
    {
        return instance;
    }

    @Autowired
    private RestTemplate restTemplate;

    @Override
    public HttpRestResult postJson(String url, Object params, String token)
    {
        return postJson(url, params, null, token);
    }

    @Override
    public HttpRestResult postJson(String url, Object params)
    {
        return postJson(url, params, null, null);
    }


    @Override
    public HttpRestResult postJson(String url, Object params, HttpHeaders headers)
    {
        return postJson(url, params, headers, null);
    }

    @Override
    public HttpRestResult postJson(String url, Object params, HttpHeaders headers, String token)
    {
        ResponseEntity<String> responseEntity = null;
        String sequenceId = getSequenceId();
        try
        {
            String body = HttpUtils.getBody(params);
            if (headers == null)
            {
                headers = new HttpHeaders();
            }
            MediaType type = MediaType.APPLICATION_JSON_UTF8;
            headers.setContentType(type);
            headers.add("Accept", MediaType.APPLICATION_JSON.toString());
            if (StringUtils.isNotEmpty(token))
            {
                headers.set("token", token);
            }
//            LOG.info(BSModule.HTTP_FRAMEWORK, "send rest http request, url = %s, sequenceId = %s, headers = %s, body = %s", url, sequenceId, headers, body);
            HttpEntity<String> formEntity = new HttpEntity<String>(body, headers);
            responseEntity = restTemplate.postForEntity(url, formEntity, String.class);
            if (this.isSuccess(responseEntity))
            {
                logger.info("receive rest http response, url = %s, sequenceId = %s, token = %s, responseEntity = %s", url, sequenceId, token, responseEntity);
            } else
            {
                logger.info(
                        "receive rest http response, url = %s, sequenceId = %s, token = %s, responseEntity = %s", url, sequenceId, token, responseEntity);
            }
            return new HttpRestResult(responseEntity.getStatusCode(), responseEntity.getHeaders(), responseEntity.getBody());
        } catch (Throwable e)
        {
            String result = null;
            HttpHeaders respHeaders = null;
            int status = HttpStatus.BAD_REQUEST.value();
            if (e instanceof UnknownHttpStatusCodeException)
            {
                UnknownHttpStatusCodeException ex = (UnknownHttpStatusCodeException) e;
                result = ex.getResponseBodyAsString();
                respHeaders = ex.getResponseHeaders();
                status = ex.getRawStatusCode();
            } else if (e instanceof HttpClientErrorException)
            {
                HttpClientErrorException ex = (HttpClientErrorException) e;
                result = ex.getResponseBodyAsString();
                respHeaders = ex.getResponseHeaders();
                status = ex.getRawStatusCode();
            }
            logger.warn(String.format("receive rest http response, url = %s, sequenceId = %s, params = %s, headers = %s, token = %s, result = %s",
                    url, sequenceId, JSONUtil.toJsonString(params), headers, token, result));
//            LOG.warning(BSModule.HTTP_FRAMEWORK, e, "receive rest http response, url = %s, sequenceId = %s, params = %s, headers = %s, token = %s, result = %s",
//                    url, sequenceId, JsonUtil.toJsonString(params), headers, token, result);
            return new HttpRestResult(HttpStatus.valueOf(status), respHeaders, result);
        }
    }

    private boolean isSuccess(ResponseEntity<String> responseEntity)
    {
        if (responseEntity == null || responseEntity.getStatusCode() != HttpStatus.OK)
        {
            return false;
        }
        return true;
    }

    private String getSequenceId()
    {
//        try
//        {
//            if (stack == null || stack.getCurrentCommand() == null)
//            {
//                return null;
//            }
//            return stack.getCurrentCommand().getId();
//        } catch (Exception e)
//        {
//        }
        return UUIDUtil.uuid2();
    }

    @Override
    public boolean isSuccess(HttpRestResult result)
    {
        if (result == null || result.getHttpStatus() != HttpStatus.OK)
        {
            return false;
        }
        return true;
    }

    @Override
    public HttpRestResult sendGET(String url, Map<String, ?> uriVariables)
    {
        return sendGET(url, uriVariables, null);
    }

    @Override
    public HttpRestResult sendGET(String url)
    {
        return sendGET(url, null, null);
    }

    @Override
    public HttpRestResult sendGET(String url, Map<String, ?> uriVariables, HttpHeaders headers)
    {
        ResponseEntity<String> responseEntity = null;
        try
        {
            responseEntity = restTemplate.getForEntity(url, String.class, Collections.emptyMap());
            if (this.isSuccess(responseEntity))
            {
                logger.info(
                        "receive rest http response, url = %s, params = %s, headers = %s, responseEntity = %s",
                        url, toString(uriVariables), headers, responseEntity);
            } else
            {
                logger.info(
                        "receive rest http response, url = %s, params = %s, headers = %s, responseEntity = %s",
                        url, toString(uriVariables), headers, responseEntity);
            }
            return new HttpRestResult(responseEntity.getStatusCode(), responseEntity.getHeaders(),
                    responseEntity.getBody());
        } catch (Exception e)
        {
            logger.info(
                    "receive rest http response, url = %s, params = %s, headers = %s, responseEntity = %s",
                    url, toString(uriVariables), headers, responseEntity);
            if (e instanceof UnknownHttpStatusCodeException)
            {
                UnknownHttpStatusCodeException ex = (UnknownHttpStatusCodeException) e;
                String result = ex.getResponseBodyAsString();
                HttpHeaders respHeaders = ex.getResponseHeaders();
                int status = ex.getRawStatusCode();
                return new HttpRestResult(HttpStatus.valueOf(status), respHeaders, result);
            } else if (e instanceof HttpClientErrorException)
            {
                HttpClientErrorException ex = (HttpClientErrorException) e;
                String result = ex.getResponseBodyAsString();
                HttpHeaders respHeaders = ex.getResponseHeaders();
                int status = ex.getRawStatusCode();
                return new HttpRestResult(HttpStatus.valueOf(status), respHeaders, result);
            }
        }
        return new HttpRestResult(HttpStatus.BAD_REQUEST);
    }

    private String toString(Object params)
    {
        if (params == null)
        {
            return null;
        }
        return JSONUtil.toJsonString(params);
    }

}
