package com.site.mountain.dao.test2;

import java.util.List;
import com.site.mountain.entity.SysDept;

public interface SysDeptDao {
    int insert(SysDept pojo);
    int insertSelective(List<SysDept> pojo);
    List<String> findList(SysDept pojo);
    int delete(SysDept pojo);

}