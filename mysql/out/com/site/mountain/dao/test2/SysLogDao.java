package com.site.mountain.dao.test2;

import java.util.List;
import com.site.mountain.entity.SysLog;

public interface SysLogDao {
    int insert(SysLog pojo);
    int insertSelective(List<SysLog> pojo);
    List<String> findList(SysLog pojo);
    int delete(SysLog pojo);

}