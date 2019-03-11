package com.site.mountain.entity;
import java.math.BigInteger;
import java.lang.String;
import java.lang.Integer;


public class SysDict {
    private BigInteger id;
    private String name;
    private String type;
    private String code;
    private String value;
    private Integer orderNum;
    private String remark;
    private Integer delFlag;

    public BigInteger getId() {
        return id;
    }
    public String getName() {
        return name;
    }
    public String getType() {
        return type;
    }
    public String getCode() {
        return code;
    }
    public String getValue() {
        return value;
    }
    public Integer getOrderNum() {
        return orderNum;
    }
    public String getRemark() {
        return remark;
    }
    public Integer getDelFlag() {
        return delFlag;
    }

    public void setId(BigInteger id) {
        this.id = id;
    }
    public void setName(String name) {
        this.name = name;
    }
    public void setType(String type) {
        this.type = type;
    }
    public void setCode(String code) {
        this.code = code;
    }
    public void setValue(String value) {
        this.value = value;
    }
    public void setOrderNum(Integer orderNum) {
        this.orderNum = orderNum;
    }
    public void setRemark(String remark) {
        this.remark = remark;
    }
    public void setDelFlag(Integer delFlag) {
        this.delFlag = delFlag;
    }

}
