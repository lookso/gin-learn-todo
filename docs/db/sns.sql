CREATE TABLE `sns_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户自增ID',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户姓名',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '用户密码',
  `phone` int(11) NOT NULL DEFAULT '0' COMMENT '手机号码',
  `nickname` varchar(128) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '小头像',
  `medium` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '大头像',
  `age` int(11) NOT NULL DEFAULT '0' COMMENT '年龄',
  `sex` tinyint(3) NOT NULL DEFAULT '0' COMMENT '1男2女',
  `astrology` varchar(32) NOT NULL DEFAULT '' COMMENT '星座',
  `status` smallint(4) NOT NULL DEFAULT '0' COMMENT '1有效 2无效',
  `source` smallint(4) NOT NULL DEFAULT '0' COMMENT '来源 0.真实用户 1.糗事百科 2.段子手',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `sns_user_community_post` (
  `id` int(11) NOT NULL DEFAULT '0' COMMENT '帖子ID',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `post_num` int(11) NOT NULL DEFAULT '0' COMMENT '帖子数',
  `comment_num` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
  `report_num` int(11) NOT NULL DEFAULT '0' COMMENT '举报数',
  `favorite_num` int(11) NOT NULL DEFAULT '0' COMMENT '收藏数',
  `zan_num` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
  `views_num` int(11) NOT NULL DEFAULT '0' COMMENT '浏览数',
  `status` smallint(4) NOT NULL DEFAULT '1' COMMENT '1:默认 1:自动审核上线,2:运营审核上线 101:运营审核下线,201:用户主动删除 202:运营审核删除',
  `status_top` tinyint(3) NOT NULL DEFAULT '0' COMMENT '置顶状态 0:默认状态，1:置顶',
  `allow_comment` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否允许评论0允许1不允许',
  `status_hot` tinyint(3) NOT NULL DEFAULT '0' COMMENT '热度状态 0:默认状态,1:热点',
  `source` tinyint(3) NOT NULL DEFAULT '0' COMMENT '0.真实用户发帖1糗事百科 2.段子手app',
  `content_type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '内容类型:1.段子 2.图片 3.gif图 4.视频,5.热门',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '帖子标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '帖子内容',
  `source_created_at` int(11) NOT NULL COMMENT '内容数据源创建时间',
  `source_published_at` int(11) NOT NULL COMMENT '内容数据源发布时间',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `publish_time` int(11) NOT NULL DEFAULT '0' COMMENT '发布时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区用户帖子表';


CREATE TABLE `alarms` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `main_title` varchar(45) NOT NULL DEFAULT '' COMMENT '主标题',
  `sub_title` varchar(45) NOT NULL DEFAULT '' COMMENT '副标题',
  `cover_img` varchar(255) NOT NULL DEFAULT '' COMMENT 'cover图片',
  `unlock_type` varchar(32) NOT NULL DEFAULT '' COMMENT '解锁类型:免费:free 快乐币解锁:coin 快乐能量解锁:energy 视频广告解锁 ad',
  `pay_money` int(11) NOT NULL DEFAULT '0' COMMENT '付费金额',
  `tag` varchar(32) NOT NULL DEFAULT '' COMMENT 'new 最新,hot 最热',
  `set_play_num` int(11) NOT NULL DEFAULT '0' COMMENT '后台设置的播放量',
  `real_play_num` int(11) NOT NULL DEFAULT '0' COMMENT '真实播放量',
  `duration` int(10) NOT NULL DEFAULT '0' COMMENT '音频时长',
  `alarm_url` varchar(255) NOT NULL DEFAULT '' COMMENT '闹铃地址',
  `star_id` int(10) NOT NULL COMMENT '明星id',
  `start_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '开始时间',
  `status` tinyint(3) NOT NULL DEFAULT '1' COMMENT '1上线 2下线',
  `alarm_type` varchar(32) NOT NULL DEFAULT '' COMMENT '闹铃类型音频:audio,视频:video',
  `is_top` tinyint(3) NOT NULL COMMENT '是否置顶0否1是',
  `top_sort` int(10) NOT NULL COMMENT '置顶顺序',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='闹铃';

CREATE TABLE `stars` (
   `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
   `name` varchar(45) NOT NULL DEFAULT '' COMMENT '合集名称',
   `sort` int(11) NOT NULL COMMENT '按优先级升序',
   `cover_img` varchar(255) NOT NULL DEFAULT '' COMMENT '明星cover图',
   `version` varchar(45) NOT NULL DEFAULT '0.0.0' COMMENT '版本号',
   `status` tinyint(3) NOT NULL DEFAULT '1' COMMENT '1 上线线 2下线',
   `is_index` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否在首页0否1是',
   `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
   `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='明星表';


CREATE TABLE `user_alarms` (
   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
   `user_id` int(10) unsigned NOT NULL COMMENT '用户id',
   `alarm_id` int(10) unsigned NOT NULL COMMENT '闹铃id',
   `star_id` int(10) unsigned NOT NULL COMMENT '明星Id',
   `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
   `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY `uniq_id_alarm` (`user_id`,`alarm_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户解锁的闹铃';