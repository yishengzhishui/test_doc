-- 查询表数据大小
SELECT database,
       table                                                                       AS `表名`,
       sum(rows)                                                                   AS `总行数`,
       formatReadableSize(sum(data_uncompressed_bytes))                            AS `原始大小`,
       formatReadableSize(sum(data_compressed_bytes))                              AS `压缩大小`,
       round((sum(data_compressed_bytes) / sum(data_uncompressed_bytes)) * 100, 0) AS `压缩率`
FROM system.parts
where database = 'ods'
GROUP BY table, database
order by sum(data_uncompressed_bytes) desc;

--match, extract, count(distinct user_id)= uniqExact()
--使用正则匹配 extract(lower(name), 'cid([0-9]+)') as content_id
select language,
       count(distinct udid)                        as `总设备`,
       count(distinct user_id)                     as `总人数`,
       uniqExact(user_id)                          as `总人数`,
       count(*)                                    as `总订单数`,
       sum(toInt32(extract(product_id, '(\\d+)'))) as `总金币数`
from ads.daily_orders_info
where match(product_id, 'coins')
  and toDate(created_at) >= '2021-10-01'
  and toDate(created_at) < '2021-11-01'
group by language
order by language desc;

-- visitParamExtractRaw 返回字段的值，包含空格符
--visitParamExtractString 使用双引号解析字符串 返回解析后的值
select id,
       created_at,
       user_id,
       order_id,
       product_id,
       area,
       client_info_str,
       visitParamExtractString(client_info_str, '_udid')     as udid,
       visitParamExtractString(client_info_str, '_language') as language,
       visitParamExtractString(client_info_str, '_package')  as package,
       visitParamExtractString(client_info_str, '_v')        as version
from ads.daily_orders_info
where product_id is not null
  and toDate(created_at) = '${date}';


--留存 bitmapContains(a.retentions, 1) 返回 0 or 1
select a.date                                                                                  as `日期`,
       count(*)                                                                                as `新增用户数`,
       if(sum(bitmapContains(a.retentions, 1)) > 0,
          toString(round(sum(bitmapContains(a.retentions, 1)) / `新增用户数` * 100, 2)) || '%', '') as D1
from dws.daily_user_retentions as a
         right join (
    -- 查找新增设备
    select distinct user_id
    from dws.daily_user_info_with_date
    ) as b on a.user_id = b.user_id
group by a.date;

--insert into， multiIf
insert into dwd.daily_page_enter
select toDate(timestamp)                                                           as date,
       timestamp,
       ugid,
       user_id,
       app_id,
       platform_id,
       language,
       area,
       multiIf(page_source_detail like '%activity-stickers%', 'activity-stickers',
               page_source_detail like '%user-avatar-box%', 'user-avatar-box', '') as page_name
from ods.daily_client_events
where event = 'PageEnter'
  and toDate(timestamp) = '2022-03-01'
  and (page_source_detail like '%activity-stickers%' or page_source_detail like '%user-avatar-box%');


-- insert into 指定字段
insert into dwd.daily_enter_tab(date, timestamp, ugid, user_id, app_id, platform_id, language, area, page_name)
select toDate(timestamp) as date,
       timestamp,
       ugid,
       user_id,
       app_id,
       platform_id,
       language,
       area,
       page_name
from ods.daily_client_events
where event = 'EnterTab'
  and toDate(timestamp) = '2022-03-01';

-- neighbor(column, offset) 拿取下/上一行的数据
select *
from (select dd, cost, neighbor(cost, 1) as last_week_cost
      from (select toDate(toMonday(date)) as dd, sum(spend) as cost
            from ads.facebook_market_report
            group by dd
            order by dd desc))
where last_week_cost > 0;

-- argMax(arg, val) 最大 val的arg值
-- sumIf(column, cond), countIf(cond), avgIf(x, cond),
-- uniqExactIf(column, cond), argMinIf(arg, val, cond)
select a.language                                    as language,
       a.type                                        as content_type,
       a.id                                          as content_id,
       uniqExact(b.user_id)                          as coin_consume_user_count,
       sumIf(toInt32(b.coins), b.type in ('1', '2')) as coin_unlock_coins,
       countIf(b.type = '3')                         as coin_tips_count,
       uniqExactIf(b.user_id, b.type = '3')          as coin_tips_user_count
from mysql_data.contents_all as a
         global
         inner join mysql_data.coins_history_all as b on (a.language = b.language and a.id = b.content_id)
group by a.language, a.type, a.id
limit 10;

-- arrayJoin()单行转为多行(数组展开),
-- arrayDistinct()数组元素去重,
-- arrayFlatten(array_of_arrays)多维数组变一维数组
select arrayJoin(arrayDistinct(arrayFlatten(groupArray(tables)))) as tt
from system.query_log
where event_date = '2022-06-15'
  and current_database = 'dim'
  and query_kind = 'Select';

-- 累计数据计算
-- arrayEnumerate 返回[1, 2, 3, …, length (arr) ]
-- arraySlice(array, offset[, length]) 数组下标从1开始，返回一个子数组，包含从指定位置的指定长度的元素
-- arraySum([func,] arr) 返回源数组中元素的总和
-- array join https://clickhouse.com/docs/zh/sql-reference/statements/select/array-join/

select date                                   as `日期`,
       'grade_' || toString(grade)            as `等级`,
       arraySum(arraySlice(count_list, 1, i)) as `累计签约作品数`
from (
      select groupArray(date)           as dates,
             groupArray(count)          as count_list,
             groupArray(contract_grade) as grades
      from (
            select count() as count, contract_grade, toDate(contract_begin_time) as date
            from dim.daily_contribution_contents
            where contract_type != 0
            group by date, contract_grade
            order by date
               )
      group by contract_grade
         )
    array join
     dates as date, --使用别名，数组元素可以通过此别名访问，但数组本身则通过原始名称访问
     grades as grade,
     arrayEnumerate(count_list) as i
order by date, `等级`;

-- 日期函数 toStartOfMonth, formatDateTime
select toDate('2022-07-06'),
       toStartOfMonth(toDate('2022-07-06')),
       formatDateTime(toDate('2022-07-06'), '%Y-%m-01');

-- groupBitmapMerge(参数是聚合中间状态(state)) 返回merge后的数据结果值，
-- groupBitmapMergeState(参数是聚合中间状态(state)) 返回merge后的中间状态，
select toStartOfMonth(date)                     as month,
       content_id,
       groupBitmapMerge(watch_device_list)      as device_count,
       groupBitmapMergeState(watch_device_list) as device_count_state
from dws.daily_content_episode_counts
where episode_weight = 2
  and language in (5)
  and content_type in (2)
  and content_id in (399015)
group by month, content_id
order by month;

-- 指定数据拆分的时候
-- maxIf()可以将指定多行变成单行
select content_id,
       maxIf(read_bmp, episode_weight = 2)  as count_2,
       maxIf(read_bmp, episode_weight = 20) as count_20,
       round((count_20 / count_2) * 100)    as r20_r2
from (select content_id,
             episode_weight,
             groupBitmapMerge(watch_device_list) as read_bmp
      from dws.daily_content_episode_counts
      where language = 9
        and episode_weight in (2, 20)
      group by content_id, episode_weight)
group by content_id;

-- splitByChar()字符串拆分成数组
select id,
       language,
       arrayMap(x->toUInt32OrZero(x),
                splitByChar(',', also_watched_content_ids_str_recent_7_days)) as dd
from dim.daily_contents_also_watched_also_watched;


-- modulo(a, b), a % b 取余数
select ugid,
       bitmapToArray(content_ids_bmp)[modulo(100, bitmapCardinality(content_ids_bmp)) + 1] as similar_content_id
from ads.ugid_read_similar_content_ids
where language = 'id'
  and app_id = 2
  and bitmapCardinality(content_ids_bmp) != 0
limit 10;

-- 日期操作, 日期可以相减
select toStartOfMonth(subtractMonths(today(), 1)) as last_month,
       toStartOfMonth(subtractMonths(today(), 2)) as last_two_month,
       toStartOfMonth(subtractMonths(today(), 3)) as last_three_month,
       today() - toDate('2022-07-05');


-- toDecimal32(price,2) 2为精度
-- position(haystack, needle)
select date,
       multiIf(
               toUInt8(position('palmaxlimited' in lower(campaign_name))), 'palmaxlimited',
               toUInt8(position('gatherone' in campaign_name)), 'gatherone',
               toUInt8(position('findmobi' in campaign_name)), 'highmobi',
               toUInt8(position('highmobi' in campaign_name)), 'highmobi',
               toUInt8(position('tec-do' in campaign_name)), 'mobisummer',
               toUInt8(position('mobisummer' in campaign_name)), 'mobisummer',
               toUInt8(position('huntmobi' in campaign_name)), 'huntmobi', null) as sub_channel
from ads.facebook_market_report
where date = '2022-01-01';


-- rename table
-- alter table
-- optimize table final
-- drop table
alter table ads.crowd_type
    drop partition '2022-07-15';
alter
table
ods.daily_send_logs
update is_retention_on_next_day = 1
where toDate(timestamp) = '2022-03-07';
rename table ads.monthly_cartoon_read to ads.monthly_cartoon_read_episodes;
optimize table dws.daily_content_counts final; --立即执行
drop table ads.udid_media_source_date;

-- row_number() 窗口函数 指定区分和排序依据
select *, row_number() OVER (PARTITION BY (date, type, id) ORDER BY date DESC) as rn
from ads.daily_orders_info
where date = '2022-02-01'
    SETTINGS allow_experimental_window_functions = 1;

select *, row_number() OVER (PARTITION BY udid_md5 ORDER BY push_message_weight DESC) as rn
from push_queue_weight_prepare_20220606;


-- with
-- with 作为常量
with (select groupBitmapState(user_id)
      from dws.daily_user_info_with_date
      where profile_is_robot = 1) as data_temp
select date, bitmapOrCardinality(data_temp, users)
from (select date, groupBitmapMergeState(comment_user_list) as users
      from dws.daily_content_counts
      where date >= '2022-02-28'
        and date <= '2022-03-07'
        and language = '1'
      group by date) as a;

-- 一个数据集合(临时表)
with TValidReadContentId as (
select ugid, round(divide(max_episode_weight, open_episodes_count) * 100) as process
from ugid_read_progress_20220530
where app_id = 1
  and language = 4
  and max_episode_weight > 5
  and content_id in (10)
    )
select *
from udid_push_info_final_20220530 T
where `app_id` = 1
  and `language` = 4
  and (ugid in (select ugid from TValidReadContentId));


-- interval link: https://clickhouse.com/docs/zh/sql-reference/data-types/special-data-types/interval/
-- 日期时间操作相关，与add/subtract year/month 相似
select (today() + toIntervalDay(1)) as a, addDays(today(), 1) as b, if(a = b, '=', '!=') as result;
select (today() - toIntervalDay(1)) as a, subtractDays(today(), 1) as b, if(a = b, '=', '!=') as result;


-- bitmap 位图
-- groupBitmapMerge 将中间聚合状态作为参数，组合状态以完成聚合
-- groupBitmapMergeState 将中间聚合状态作为参数，返回中间聚合状态
-- groupBitmapState 返回中间聚合状态
-- groupBitmapStateIf(bitmap,条件) 返回符合条件的bitmap中间状态

-- create table
-- MergeTree
create table if not exists ads.sdk_ads_report_temp
(
    id               UInt64,
    date             Date,
    app              Nullable(UInt8),
    placement        Nullable(String),
    platform         Nullable(UInt8),
    country          Nullable(String),
    impressions      Nullable(UInt64),
    clicks           Nullable(UInt64),
    revenue          Decimal(12, 4),
    ad_type          Nullable(UInt8),
    requests         Nullable(UInt64),
    matched_requests Nullable(UInt64),
    format           Nullable(UInt8),
    vendor           String
)
    engine = MergeTree PARTITION BY date ORDER BY (date, vendor, id)
        SETTINGS index_granularity = 8192;

-- ReplacingMergeTree 可以用 final
create table if not exists dws.device_favorite_for_fcm_push
(
    push_token_id Int64,
    content_id    Int32,
    language      String,
    is_deleted    Int8

)
    engine = ReplacingMergeTree PARTITION BY language ORDER BY (language, push_token_id, content_id)
        SETTINGS index_granularity = 8192;

-- AggregatingMergeTree，
-- SimpleAggregateFunction 用于数据汇总后进行聚合
-- link： https://clickhouse.com/docs/zh/sql-reference/data-types/simpleaggregatefunction/
CREATE TABLE if not exists ads.daily_contribution_content_count
(
    date                                    Date COMMENT '日期',
    language                                Int32 COMMENT '语种',
    content_type                            Int32 COMMENT '作品类型',
    contribution_count                      SimpleAggregateFunction(anyLast, Nullable(Int32)) COMMENT '投稿作品数',
    content_online_count                    SimpleAggregateFunction(anyLast, Nullable(Int32)) COMMENT '上线作品数',
    content_update_count_except_new_content SimpleAggregateFunction(anyLast, Nullable(Int32)) COMMENT '有更新的作品数(除新作)',
    contribution_episode_count              SimpleAggregateFunction(anyLast, Nullable(Int32)) COMMENT '投稿章节数',
    episode_online_count                    SimpleAggregateFunction(anyLast, Nullable(Int32)) COMMENT '上线章节数',
    episode_update_count_except_new_content SimpleAggregateFunction(anyLast, Nullable(Int32)) COMMENT '上线章节数(除新作)'
)
    ENGINE = AggregatingMergeTree()
        PARTITION BY date
        ORDER BY (date, language, content_type);

-- TTL now() + INTERVAL 3 DAY 表中的内容 三天后失效
-- TTL toDate(timestamp) + toIntervalDay(3); timestamp是表中字段,列数据过期

-- create table as
CREATE TABLE if not exists dim.daily_contribution_contents_20220721
    as dim.daily_contribution_contents;

--startsWith(s, 前缀)以指定的前缀开头，则返回1，否则返回0
--toFixedString(s,N),将String类型的参数转换为FixedString(N)类型的值（具有固定长度N的字符串）
-- FixedString: https://clickhouse.com/docs/zh/sql-reference/data-types/fixedstring/
-- 二进制表示的哈希值（MD5使用FixedString(16)
select *
from temp.push_udid_ugid_map_20220726
where udid_md5 = toFixedString(MD5('83f5f3f0-5c2f-4e66-bf6d-403c8f75fedd'), 16);

--udid的MD5, unhex是hex的反操作
select toFixedString(if(match(udid, '^[a-fA-F0-9]{32}$'), unhex(udid), MD5(udid)), 16)
from (select distinct udid
      from manhuagu_ugid.udids)
;

--substring(s,offset,length),substr(s,offset,length)
--以字节为单位截取指定位置字符串，返回以’offset’位置为开头，长度为’length’的子串。’offset’从1开始（与标准SQL相同）。’offset’和’length’参数必须是常量。
-- trimBoth(string)删除两侧的空白字符
-- 字符串函数 link： https://clickhouse.com/docs/zh/sql-reference/functions/string-functions/#trimboth

--between 8 and 9 == >=8 and <=9
-- coalesce(x,...) 任何数量的非复合类型的参数。所有参数的数据类型必须兼容
-- 返回值 第一个非’NULL`参数；如果所有参数都是’NULL`，那就返回null。






