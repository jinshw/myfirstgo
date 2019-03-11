package com.site.mountain.dao.test2;

import java.util.List;
import com.site.mountain.entity.SysMenu;

public interface SysMenuDao {
    int insert(SysMenu pojo);
    int insertSelective(List<SysMenu> pojo);
    List<String> findList(SysMenu pojo);
    int delete(SysMenu pojo);

}