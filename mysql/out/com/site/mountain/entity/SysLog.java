package com.site.mountain.entity;
import java.math.BigInteger;
import java.lang.String;
import java.sql.Timestamp;


public class SysLog {
    private BigInteger id;
    private String username;
    private String operation;
    private String method;
    private String params;
    private BigInteger time;
    private String ip;
    private Timestamp createDate;

    public BigInteger getId() {
        return id;
    }
    public String getUsername() {
        return username;
    }
    public String getOperation() {
        return operation;
    }
    public String getMethod() {
        return method;
    }
    public String getParams() {
        return params;
    }
    public BigInteger getTime() {
        return time;
    }
    public String getIp() {
        return ip;
    }
    public Timestamp getCreateDate() {
        return createDate;
    }

    public void setId(BigInteger id) {
        this.id = id;
    }
    public void setUsername(String username) {
        this.username = username;
    }
    public void setOperation(String operation) {
        this.operation = operation;
    }
    public void setMethod(String method) {
        this.method = method;
    }
    public void setParams(String params) {
        this.params = params;
    }
    public void setTime(BigInteger time) {
        this.time = time;
    }
    public void setIp(String ip) {
        this.ip = ip;
    }
    public void setCreateDate(Timestamp createDate) {
        this.createDate = createDate;
    }

}
