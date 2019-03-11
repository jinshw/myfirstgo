package com.site.mountain.dao.test2;

import java.util.List;
import com.site.mountain.entity.SysDict;

public interface SysDictDao {
    int insert(SysDict pojo);
    int insertSelective(List<SysDict> pojo);
    List<String> findList(SysDict pojo);
    int delete(SysDict pojo);

}