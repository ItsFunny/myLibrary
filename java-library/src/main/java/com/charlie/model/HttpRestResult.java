package com.charlie.model;

import lombok.Data;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;

@Data
public class HttpRestResult
{
	private HttpStatus httpStatus;
	private HttpHeaders headers;
	private String result;

	public HttpRestResult() {
		super();
	}

	public HttpRestResult(HttpStatus httpStatus) {
		super();
		this.httpStatus = httpStatus;
	}

	public HttpRestResult(HttpStatus httpStatus, HttpHeaders headers, String result) {
		super();
		this.httpStatus = httpStatus;
		this.headers = headers;
		this.result = result;
	}

	public HttpStatus getHttpStatus() {
		return httpStatus;
	}

	public void setHttpStatus(HttpStatus httpStatus) {
		this.httpStatus = httpStatus;
	}

	public HttpHeaders getHeaders() {
		return headers;
	}

	public void setHeaders(HttpHeaders headers) {
		this.headers = headers;
	}

	public String getResult() {
		return result;
	}

	public void setResult(String result) {
		this.result = result;
	}

	@Override
	public String toString() {
		return "HttpRestResult [httpStatus=" + httpStatus + ", headers=" + headers + ", result=" + result + "]";
	}

}
