package com.sykj.river.master.service.impl;

import com.sykj.river.master.model.River;
import com.sykj.river.master.model.RiverExample;
import com.sykj.river.master.persistence.RiverMapper;
import com.sykj.river.master.service.RiverService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * description: RiverServiceImpl
 *
 * @generator com.github.wangxianzhuo.template-gen
 * @date 2020-05-21T14:30:50+08:00
 * @version: 1.0
 */
@Service
public class RiverServiceImpl implements RiverService {

    @Autowired
    private RiverMapper riverMapper;

    @Override
    public River get(String id) {
        return riverMapper.selectByPrimaryKey(id);
    }

    @Override
    public List<River> all() {
        return riverMapper.selectByExample(new RiverExample());
    }

    @Override
    public void add(River river) {
        riverMapper.insertSelective(river);
    }

    @Override
    public void modify(River river) {
        riverMapper.updateByPrimaryKeySelective(river);
    }

    @Override
    public void remove(String id) {
        riverMapper.deleteByPrimaryKey(id);
    }
}
