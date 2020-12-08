package com.charlie.service;

import com.charlie.model.HttpRestResult;
import org.springframework.http.HttpHeaders;

import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-11-22 15:57
 */
public interface IHttpRestService
{
    HttpRestResult postJson(String url, Object params, String token);

    HttpRestResult postJson(String url, Object params, HttpHeaders headers);

    HttpRestResult postJson(String url, Object params, HttpHeaders headers, String token);

    HttpRestResult postJson(String url, Object params);

    boolean isSuccess(HttpRestResult result);

    HttpRestResult sendGET(String url);

    HttpRestResult sendGET(String url, Map<String, ?> uriVariables);

    HttpRestResult sendGET(String url, Map<String, ?> uriVariables, HttpHeaders headers);
}
