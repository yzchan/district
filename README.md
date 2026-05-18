# 中国行政区划知识库

这是一个面向中国行政区划的中文知识型开源项目，目标是把政策依据、官方来源、历史沿革和数据快照整理成一套可追溯、可引用、可被大模型直接消费的知识底座。

`data/distict.json` 是当前核心数据文件，下载时间为 `2026-05-18`。

## 政策背景

- 民政部行政区划代码栏目：[https://www.mca.gov.cn/n156/n186/index.html](https://www.mca.gov.cn/n156/n186/index.html)
- 国家地名信息库：[https://dmfw.mca.gov.cn/](https://dmfw.mca.gov.cn/)
- 国家地名信息库接口说明：[https://dmfw.mca.gov.cn/interface.html](https://dmfw.mca.gov.cn/interface.html)
- 国家统计局《统计用区划代码和城乡划分代码编制规则》：[https://www.stats.gov.cn/sj/tjbz/gjtjbz/202302/t20230213_1902741.html](https://www.stats.gov.cn/sj/tjbz/gjtjbz/202302/t20230213_1902741.html)
- 国家统计局咨询公开“请问最新的统计用区划代码和城乡划分代码在哪里查看？”：[https://www.stats.gov.cn/hd/lyzx/zxgk/202509/t20250903_1960996.html](https://www.stats.gov.cn/hd/lyzx/zxgk/202509/t20250903_1960996.html)

民政部口径是行政区划建制信息主线；统计局口径属于统计标准体系；第三方平台数据仅作兼容参考。

## 项目愿景

让中国行政区划代码的来源更清楚、历史更完整、口径更统一，并能稳定支撑 JS、SQL、CSV、JSON 和说明文档的生成。

## 对 AI 的要求

AI 必须在明确来源、明确时点、明确口径的前提下工作，区分行政区划建制、统计区划体系和兼容性平台数据，避免把区划代码当作静态常量。
