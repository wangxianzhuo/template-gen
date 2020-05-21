package com.sykj.river.master.service;

import com.sykj.river.master.model.River;

import java.util.List;

/**
 * RiverService
 *
 * @generator com.github.wangxianzhuo.template-gen
 * @date 2020-05-21T14:30:50+08:00
 */
public interface RiverService {
    River get(String id);

    List<River> all();

    void add(River river);

    void modify(River river);

    void remove(String id);
}
