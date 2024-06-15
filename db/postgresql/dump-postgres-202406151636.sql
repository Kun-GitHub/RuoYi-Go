--
-- PostgreSQL database cluster dump
--

-- Started on 2024-06-15 16:36:51

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS;

--
-- User Configurations
--






--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.7 (Debian 14.7-1.pgdg110+1)
-- Dumped by pg_dump version 15.3

-- Started on 2024-06-15 16:36:51

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3306 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2024-06-15 16:36:51

--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

\connect postgres

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.7 (Debian 14.7-1.pgdg110+1)
-- Dumped by pg_dump version 15.3

-- Started on 2024-06-15 16:36:51

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 24577)
-- Name: sys_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_config (
    config_id integer NOT NULL,
    config_name character varying(100),
    config_key character varying(100),
    config_value character varying(500),
    config_type character(1),
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone,
    remark character varying(500)
);


ALTER TABLE public.sys_config OWNER TO postgres;

--
-- TOC entry 3432 (class 0 OID 0)
-- Dependencies: 209
-- Name: TABLE sys_config; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_config IS '参数配置表';


--
-- TOC entry 3433 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.config_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.config_id IS '参数主键';


--
-- TOC entry 3434 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.config_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.config_name IS '参数名称';


--
-- TOC entry 3435 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.config_key; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.config_key IS '参数键名';


--
-- TOC entry 3436 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.config_value; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.config_value IS '参数键值';


--
-- TOC entry 3437 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.config_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.config_type IS '系统内置（Y是 N否）';


--
-- TOC entry 3438 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.create_by IS '创建者';


--
-- TOC entry 3439 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.create_time IS '创建时间';


--
-- TOC entry 3440 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.update_by IS '更新者';


--
-- TOC entry 3441 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.update_time IS '更新时间';


--
-- TOC entry 3442 (class 0 OID 0)
-- Dependencies: 209
-- Name: COLUMN sys_config.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_config.remark IS '备注';


--
-- TOC entry 210 (class 1259 OID 24582)
-- Name: sys_dept; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_dept (
    dept_id bigint NOT NULL,
    parent_id bigint,
    ancestors character varying(50),
    dept_name character varying(30),
    order_num integer,
    leader character varying(20),
    phone character varying(11),
    email character varying(50),
    status character(1),
    del_flag character(1),
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone
);


ALTER TABLE public.sys_dept OWNER TO postgres;

--
-- TOC entry 3443 (class 0 OID 0)
-- Dependencies: 210
-- Name: TABLE sys_dept; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_dept IS '部门表';


--
-- TOC entry 3444 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.dept_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.dept_id IS '部门id';


--
-- TOC entry 3445 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.parent_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.parent_id IS '父部门id';


--
-- TOC entry 3446 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.ancestors; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.ancestors IS '祖级列表';


--
-- TOC entry 3447 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.dept_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.dept_name IS '部门名称';


--
-- TOC entry 3448 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.order_num; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.order_num IS '显示顺序';


--
-- TOC entry 3449 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.leader; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.leader IS '负责人';


--
-- TOC entry 3450 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.phone; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.phone IS '联系电话';


--
-- TOC entry 3451 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.email; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.email IS '邮箱';


--
-- TOC entry 3452 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.status IS '部门状态（0正常 1停用）';


--
-- TOC entry 3453 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.del_flag; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.del_flag IS '删除标志（0代表存在 2代表删除）';


--
-- TOC entry 3454 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.create_by IS '创建者';


--
-- TOC entry 3455 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.create_time IS '创建时间';


--
-- TOC entry 3456 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.update_by IS '更新者';


--
-- TOC entry 3457 (class 0 OID 0)
-- Dependencies: 210
-- Name: COLUMN sys_dept.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dept.update_time IS '更新时间';


--
-- TOC entry 211 (class 1259 OID 24585)
-- Name: sys_dict_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_dict_data (
    dict_code bigint NOT NULL,
    dict_sort integer,
    dict_label character varying(100),
    dict_value character varying(100),
    dict_type character varying(100),
    css_class character varying(100),
    list_class character varying(100),
    is_default character(1),
    status character(1),
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone,
    remark character varying(500)
);


ALTER TABLE public.sys_dict_data OWNER TO postgres;

--
-- TOC entry 3458 (class 0 OID 0)
-- Dependencies: 211
-- Name: TABLE sys_dict_data; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_dict_data IS '字典数据表';


--
-- TOC entry 3459 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.dict_code; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.dict_code IS '字典编码';


--
-- TOC entry 3460 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.dict_sort; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.dict_sort IS '字典排序';


--
-- TOC entry 3461 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.dict_label; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.dict_label IS '字典标签';


--
-- TOC entry 3462 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.dict_value; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.dict_value IS '字典键值';


--
-- TOC entry 3463 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.dict_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.dict_type IS '字典类型';


--
-- TOC entry 3464 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.css_class; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.css_class IS '样式属性（其他样式扩展）';


--
-- TOC entry 3465 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.list_class; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.list_class IS '表格回显样式';


--
-- TOC entry 3466 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.is_default; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.is_default IS '是否默认（Y是 N否）';


--
-- TOC entry 3467 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.status IS '状态（0正常 1停用）';


--
-- TOC entry 3468 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.create_by IS '创建者';


--
-- TOC entry 3469 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.create_time IS '创建时间';


--
-- TOC entry 3470 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.update_by IS '更新者';


--
-- TOC entry 3471 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.update_time IS '更新时间';


--
-- TOC entry 3472 (class 0 OID 0)
-- Dependencies: 211
-- Name: COLUMN sys_dict_data.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_data.remark IS '备注';


--
-- TOC entry 212 (class 1259 OID 24590)
-- Name: sys_dict_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_dict_type (
    dict_id bigint NOT NULL,
    dict_name character varying(100),
    dict_type character varying(100),
    status character(1),
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone,
    remark character varying(500)
);


ALTER TABLE public.sys_dict_type OWNER TO postgres;

--
-- TOC entry 3473 (class 0 OID 0)
-- Dependencies: 212
-- Name: TABLE sys_dict_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_dict_type IS '字典类型表';


--
-- TOC entry 3474 (class 0 OID 0)
-- Dependencies: 212
-- Name: COLUMN sys_dict_type.dict_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_type.dict_id IS '字典主键';


--
-- TOC entry 3475 (class 0 OID 0)
-- Dependencies: 212
-- Name: COLUMN sys_dict_type.dict_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_type.dict_name IS '字典名称';


--
-- TOC entry 3476 (class 0 OID 0)
-- Dependencies: 212
-- Name: COLUMN sys_dict_type.dict_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_type.dict_type IS '字典类型';


--
-- TOC entry 3477 (class 0 OID 0)
-- Dependencies: 212
-- Name: COLUMN sys_dict_type.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_type.status IS '状态（0正常 1停用）';


--
-- TOC entry 3478 (class 0 OID 0)
-- Dependencies: 212
-- Name: COLUMN sys_dict_type.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_type.create_by IS '创建者';


--
-- TOC entry 3479 (class 0 OID 0)
-- Dependencies: 212
-- Name: COLUMN sys_dict_type.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_type.create_time IS '创建时间';


--
-- TOC entry 3480 (class 0 OID 0)
-- Dependencies: 212
-- Name: COLUMN sys_dict_type.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_type.update_by IS '更新者';


--
-- TOC entry 3481 (class 0 OID 0)
-- Dependencies: 212
-- Name: COLUMN sys_dict_type.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_type.update_time IS '更新时间';


--
-- TOC entry 3482 (class 0 OID 0)
-- Dependencies: 212
-- Name: COLUMN sys_dict_type.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dict_type.remark IS '备注';


--
-- TOC entry 213 (class 1259 OID 24595)
-- Name: sys_job; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_job (
    job_id bigint NOT NULL,
    job_name character varying(64) NOT NULL,
    job_group character varying(64) NOT NULL,
    invoke_target character varying(500) NOT NULL,
    cron_expression character varying(255),
    misfire_policy character varying(20),
    concurrent character(1),
    status character(1),
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone,
    remark character varying(500)
);


ALTER TABLE public.sys_job OWNER TO postgres;

--
-- TOC entry 3483 (class 0 OID 0)
-- Dependencies: 213
-- Name: TABLE sys_job; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_job IS '定时任务调度表';


--
-- TOC entry 3484 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.job_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.job_id IS '任务ID';


--
-- TOC entry 3485 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.job_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.job_name IS '任务名称';


--
-- TOC entry 3486 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.job_group; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.job_group IS '任务组名';


--
-- TOC entry 3487 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.invoke_target; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.invoke_target IS '调用目标字符串';


--
-- TOC entry 3488 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.cron_expression; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.cron_expression IS 'cron执行表达式';


--
-- TOC entry 3489 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.misfire_policy; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.misfire_policy IS '计划执行错误策略（1立即执行 2执行一次 3放弃执行）';


--
-- TOC entry 3490 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.concurrent; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.concurrent IS '是否并发执行（0允许 1禁止）';


--
-- TOC entry 3491 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.status IS '状态（0正常 1暂停）';


--
-- TOC entry 3492 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.create_by IS '创建者';


--
-- TOC entry 3493 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.create_time IS '创建时间';


--
-- TOC entry 3494 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.update_by IS '更新者';


--
-- TOC entry 3495 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.update_time IS '更新时间';


--
-- TOC entry 3496 (class 0 OID 0)
-- Dependencies: 213
-- Name: COLUMN sys_job.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job.remark IS '备注信息';


--
-- TOC entry 214 (class 1259 OID 24600)
-- Name: sys_job_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_job_log (
    job_log_id bigint NOT NULL,
    job_name character varying(64) NOT NULL,
    job_group character varying(64) NOT NULL,
    invoke_target character varying(500) NOT NULL,
    job_message character varying(500),
    status character(1),
    exception_info character varying(2000),
    create_time timestamp without time zone
);


ALTER TABLE public.sys_job_log OWNER TO postgres;

--
-- TOC entry 3497 (class 0 OID 0)
-- Dependencies: 214
-- Name: TABLE sys_job_log; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_job_log IS '定时任务调度日志表';


--
-- TOC entry 3498 (class 0 OID 0)
-- Dependencies: 214
-- Name: COLUMN sys_job_log.job_log_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job_log.job_log_id IS '任务日志ID';


--
-- TOC entry 3499 (class 0 OID 0)
-- Dependencies: 214
-- Name: COLUMN sys_job_log.job_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job_log.job_name IS '任务名称';


--
-- TOC entry 3500 (class 0 OID 0)
-- Dependencies: 214
-- Name: COLUMN sys_job_log.job_group; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job_log.job_group IS '任务组名';


--
-- TOC entry 3501 (class 0 OID 0)
-- Dependencies: 214
-- Name: COLUMN sys_job_log.invoke_target; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job_log.invoke_target IS '调用目标字符串';


--
-- TOC entry 3502 (class 0 OID 0)
-- Dependencies: 214
-- Name: COLUMN sys_job_log.job_message; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job_log.job_message IS '日志信息';


--
-- TOC entry 3503 (class 0 OID 0)
-- Dependencies: 214
-- Name: COLUMN sys_job_log.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job_log.status IS '执行状态（0正常 1失败）';


--
-- TOC entry 3504 (class 0 OID 0)
-- Dependencies: 214
-- Name: COLUMN sys_job_log.exception_info; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job_log.exception_info IS '异常信息';


--
-- TOC entry 3505 (class 0 OID 0)
-- Dependencies: 214
-- Name: COLUMN sys_job_log.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_job_log.create_time IS '创建时间';


--
-- TOC entry 215 (class 1259 OID 24605)
-- Name: sys_logininfor; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_logininfor (
    info_id bigint NOT NULL,
    user_name character varying(50),
    ipaddr character varying(128),
    login_location character varying(255),
    browser character varying(50),
    os character varying(50),
    status character(1),
    msg character varying(255),
    login_time timestamp without time zone
);


ALTER TABLE public.sys_logininfor OWNER TO postgres;

--
-- TOC entry 3506 (class 0 OID 0)
-- Dependencies: 215
-- Name: TABLE sys_logininfor; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_logininfor IS '系统访问记录';


--
-- TOC entry 3507 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN sys_logininfor.info_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_logininfor.info_id IS '访问ID';


--
-- TOC entry 3508 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN sys_logininfor.user_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_logininfor.user_name IS '用户账号';


--
-- TOC entry 3509 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN sys_logininfor.ipaddr; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_logininfor.ipaddr IS '登录IP地址';


--
-- TOC entry 3510 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN sys_logininfor.login_location; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_logininfor.login_location IS '登录地点';


--
-- TOC entry 3511 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN sys_logininfor.browser; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_logininfor.browser IS '浏览器类型';


--
-- TOC entry 3512 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN sys_logininfor.os; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_logininfor.os IS '操作系统';


--
-- TOC entry 3513 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN sys_logininfor.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_logininfor.status IS '登录状态（0成功 1失败）';


--
-- TOC entry 3514 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN sys_logininfor.msg; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_logininfor.msg IS '提示消息';


--
-- TOC entry 3515 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN sys_logininfor.login_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_logininfor.login_time IS '访问时间';


--
-- TOC entry 216 (class 1259 OID 24610)
-- Name: sys_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_menu (
    menu_id bigint NOT NULL,
    menu_name character varying(50) NOT NULL,
    parent_id bigint,
    order_num integer,
    path character varying(200),
    component character varying(255),
    query character varying(255),
    is_frame integer,
    is_cache integer,
    menu_type character(1),
    visible character(1),
    status character(1),
    perms character varying(100),
    icon character varying(100),
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone,
    remark character varying(500)
);


ALTER TABLE public.sys_menu OWNER TO postgres;

--
-- TOC entry 3516 (class 0 OID 0)
-- Dependencies: 216
-- Name: TABLE sys_menu; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_menu IS '菜单权限表';


--
-- TOC entry 3517 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.menu_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.menu_id IS '菜单ID';


--
-- TOC entry 3518 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.menu_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.menu_name IS '菜单名称';


--
-- TOC entry 3519 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.parent_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.parent_id IS '父菜单ID';


--
-- TOC entry 3520 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.order_num; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.order_num IS '显示顺序';


--
-- TOC entry 3521 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.path IS '路由地址';


--
-- TOC entry 3522 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.component; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.component IS '组件路径';


--
-- TOC entry 3523 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.query; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.query IS '路由参数';


--
-- TOC entry 3524 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.is_frame; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.is_frame IS '是否为外链（0是 1否）';


--
-- TOC entry 3525 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.is_cache; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.is_cache IS '是否缓存（0缓存 1不缓存）';


--
-- TOC entry 3526 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.menu_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.menu_type IS '菜单类型（M目录 C菜单 F按钮）';


--
-- TOC entry 3527 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.visible; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.visible IS '菜单状态（0显示 1隐藏）';


--
-- TOC entry 3528 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.status IS '菜单状态（0正常 1停用）';


--
-- TOC entry 3529 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.perms; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.perms IS '权限标识';


--
-- TOC entry 3530 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.icon; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.icon IS '菜单图标';


--
-- TOC entry 3531 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.create_by IS '创建者';


--
-- TOC entry 3532 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.create_time IS '创建时间';


--
-- TOC entry 3533 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.update_by IS '更新者';


--
-- TOC entry 3534 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.update_time IS '更新时间';


--
-- TOC entry 3535 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN sys_menu.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_menu.remark IS '备注';


--
-- TOC entry 217 (class 1259 OID 24615)
-- Name: sys_notice; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_notice (
    notice_id integer NOT NULL,
    notice_title character varying(50) NOT NULL,
    notice_type character(1) NOT NULL,
    notice_content bytea,
    status character(1),
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone,
    remark character varying(255)
);


ALTER TABLE public.sys_notice OWNER TO postgres;

--
-- TOC entry 3536 (class 0 OID 0)
-- Dependencies: 217
-- Name: TABLE sys_notice; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_notice IS '通知公告表';


--
-- TOC entry 3537 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.notice_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.notice_id IS '公告ID';


--
-- TOC entry 3538 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.notice_title; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.notice_title IS '公告标题';


--
-- TOC entry 3539 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.notice_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.notice_type IS '公告类型（1通知 2公告）';


--
-- TOC entry 3540 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.notice_content; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.notice_content IS '公告内容';


--
-- TOC entry 3541 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.status IS '公告状态（0正常 1关闭）';


--
-- TOC entry 3542 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.create_by IS '创建者';


--
-- TOC entry 3543 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.create_time IS '创建时间';


--
-- TOC entry 3544 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.update_by IS '更新者';


--
-- TOC entry 3545 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.update_time IS '更新时间';


--
-- TOC entry 3546 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN sys_notice.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_notice.remark IS '备注';


--
-- TOC entry 218 (class 1259 OID 24620)
-- Name: sys_oper_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_oper_log (
    oper_id bigint NOT NULL,
    title character varying(50),
    business_type integer,
    method character varying(100),
    request_method character varying(10),
    operator_type integer,
    oper_name character varying(50),
    dept_name character varying(50),
    oper_url character varying(255),
    oper_ip character varying(128),
    oper_location character varying(255),
    oper_param character varying(2000),
    json_result character varying(2000),
    status integer,
    error_msg character varying(2000),
    oper_time timestamp without time zone,
    cost_time bigint
);


ALTER TABLE public.sys_oper_log OWNER TO postgres;

--
-- TOC entry 3547 (class 0 OID 0)
-- Dependencies: 218
-- Name: TABLE sys_oper_log; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_oper_log IS '操作日志记录';


--
-- TOC entry 3548 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.oper_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.oper_id IS '日志主键';


--
-- TOC entry 3549 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.title; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.title IS '模块标题';


--
-- TOC entry 3550 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.business_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.business_type IS '业务类型（0其它 1新增 2修改 3删除）';


--
-- TOC entry 3551 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.method; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.method IS '方法名称';


--
-- TOC entry 3552 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.request_method; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.request_method IS '请求方式';


--
-- TOC entry 3553 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.operator_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.operator_type IS '操作类别（0其它 1后台用户 2手机端用户）';


--
-- TOC entry 3554 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.oper_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.oper_name IS '操作人员';


--
-- TOC entry 3555 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.dept_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.dept_name IS '部门名称';


--
-- TOC entry 3556 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.oper_url; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.oper_url IS '请求URL';


--
-- TOC entry 3557 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.oper_ip; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.oper_ip IS '主机地址';


--
-- TOC entry 3558 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.oper_location; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.oper_location IS '操作地点';


--
-- TOC entry 3559 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.oper_param; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.oper_param IS '请求参数';


--
-- TOC entry 3560 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.json_result; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.json_result IS '返回参数';


--
-- TOC entry 3561 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.status IS '操作状态（0正常 1异常）';


--
-- TOC entry 3562 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.error_msg; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.error_msg IS '错误消息';


--
-- TOC entry 3563 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.oper_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.oper_time IS '操作时间';


--
-- TOC entry 3564 (class 0 OID 0)
-- Dependencies: 218
-- Name: COLUMN sys_oper_log.cost_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_oper_log.cost_time IS '消耗时间';


--
-- TOC entry 219 (class 1259 OID 24625)
-- Name: sys_post; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_post (
    post_id bigint NOT NULL,
    post_code character varying(64) NOT NULL,
    post_name character varying(50) NOT NULL,
    post_sort integer NOT NULL,
    status character(1) NOT NULL,
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone,
    remark character varying(500)
);


ALTER TABLE public.sys_post OWNER TO postgres;

--
-- TOC entry 3565 (class 0 OID 0)
-- Dependencies: 219
-- Name: TABLE sys_post; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_post IS '岗位信息表';


--
-- TOC entry 3566 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.post_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.post_id IS '岗位ID';


--
-- TOC entry 3567 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.post_code; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.post_code IS '岗位编码';


--
-- TOC entry 3568 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.post_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.post_name IS '岗位名称';


--
-- TOC entry 3569 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.post_sort; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.post_sort IS '显示顺序';


--
-- TOC entry 3570 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.status IS '状态（0正常 1停用）';


--
-- TOC entry 3571 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.create_by IS '创建者';


--
-- TOC entry 3572 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.create_time IS '创建时间';


--
-- TOC entry 3573 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.update_by IS '更新者';


--
-- TOC entry 3574 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.update_time IS '更新时间';


--
-- TOC entry 3575 (class 0 OID 0)
-- Dependencies: 219
-- Name: COLUMN sys_post.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_post.remark IS '备注';


--
-- TOC entry 220 (class 1259 OID 24630)
-- Name: sys_role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_role (
    role_id bigint NOT NULL,
    role_name character varying(30) NOT NULL,
    role_key character varying(100) NOT NULL,
    role_sort integer NOT NULL,
    data_scope character(1),
    menu_check_strictly smallint,
    dept_check_strictly smallint,
    status character(1) NOT NULL,
    del_flag character(1),
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone,
    remark character varying(500)
);


ALTER TABLE public.sys_role OWNER TO postgres;

--
-- TOC entry 3576 (class 0 OID 0)
-- Dependencies: 220
-- Name: TABLE sys_role; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_role IS '角色信息表';


--
-- TOC entry 3577 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.role_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.role_id IS '角色ID';


--
-- TOC entry 3578 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.role_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.role_name IS '角色名称';


--
-- TOC entry 3579 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.role_key; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.role_key IS '角色权限字符串';


--
-- TOC entry 3580 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.role_sort; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.role_sort IS '显示顺序';


--
-- TOC entry 3581 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.data_scope; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.data_scope IS '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）';


--
-- TOC entry 3582 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.menu_check_strictly; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.menu_check_strictly IS '菜单树选择项是否关联显示';


--
-- TOC entry 3583 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.dept_check_strictly; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.dept_check_strictly IS '部门树选择项是否关联显示';


--
-- TOC entry 3584 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.status IS '角色状态（0正常 1停用）';


--
-- TOC entry 3585 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.del_flag; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.del_flag IS '删除标志（0代表存在 2代表删除）';


--
-- TOC entry 3586 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.create_by IS '创建者';


--
-- TOC entry 3587 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.create_time IS '创建时间';


--
-- TOC entry 3588 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.update_by IS '更新者';


--
-- TOC entry 3589 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.update_time IS '更新时间';


--
-- TOC entry 3590 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN sys_role.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role.remark IS '备注';


--
-- TOC entry 221 (class 1259 OID 24635)
-- Name: sys_role_dept; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_role_dept (
    role_id bigint NOT NULL,
    dept_id bigint NOT NULL
);


ALTER TABLE public.sys_role_dept OWNER TO postgres;

--
-- TOC entry 3591 (class 0 OID 0)
-- Dependencies: 221
-- Name: TABLE sys_role_dept; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_role_dept IS '角色和部门关联表';


--
-- TOC entry 3592 (class 0 OID 0)
-- Dependencies: 221
-- Name: COLUMN sys_role_dept.role_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role_dept.role_id IS '角色ID';


--
-- TOC entry 3593 (class 0 OID 0)
-- Dependencies: 221
-- Name: COLUMN sys_role_dept.dept_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role_dept.dept_id IS '部门ID';


--
-- TOC entry 222 (class 1259 OID 24638)
-- Name: sys_role_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_role_menu (
    role_id bigint NOT NULL,
    menu_id bigint NOT NULL
);


ALTER TABLE public.sys_role_menu OWNER TO postgres;

--
-- TOC entry 3594 (class 0 OID 0)
-- Dependencies: 222
-- Name: TABLE sys_role_menu; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_role_menu IS '角色和菜单关联表';


--
-- TOC entry 3595 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN sys_role_menu.role_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role_menu.role_id IS '角色ID';


--
-- TOC entry 3596 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN sys_role_menu.menu_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_role_menu.menu_id IS '菜单ID';


--
-- TOC entry 223 (class 1259 OID 24641)
-- Name: sys_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_user (
    user_id bigint NOT NULL,
    dept_id bigint,
    user_name character varying(30) NOT NULL,
    nick_name character varying(30) NOT NULL,
    user_type character varying(2),
    email character varying(50),
    phonenumber character varying(11),
    sex character(1),
    avatar character varying(100),
    password character varying(100),
    status character(1),
    del_flag character(1),
    login_ip character varying(128),
    login_date timestamp without time zone,
    create_by character varying(64),
    create_time timestamp without time zone,
    update_by character varying(64),
    update_time timestamp without time zone,
    remark character varying(500)
);


ALTER TABLE public.sys_user OWNER TO postgres;

--
-- TOC entry 3597 (class 0 OID 0)
-- Dependencies: 223
-- Name: TABLE sys_user; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_user IS '用户信息表';


--
-- TOC entry 3598 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.user_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.user_id IS '用户ID';


--
-- TOC entry 3599 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.dept_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.dept_id IS '部门ID';


--
-- TOC entry 3600 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.user_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.user_name IS '用户账号';


--
-- TOC entry 3601 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.nick_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.nick_name IS '用户昵称';


--
-- TOC entry 3602 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.user_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.user_type IS '用户类型（00系统用户）';


--
-- TOC entry 3603 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.email; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.email IS '用户邮箱';


--
-- TOC entry 3604 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.phonenumber; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.phonenumber IS '手机号码';


--
-- TOC entry 3605 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.sex; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.sex IS '用户性别（0男 1女 2未知）';


--
-- TOC entry 3606 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.avatar; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.avatar IS '头像地址';


--
-- TOC entry 3607 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.password; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.password IS '密码';


--
-- TOC entry 3608 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.status IS '帐号状态（0正常 1停用）';


--
-- TOC entry 3609 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.del_flag; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.del_flag IS '删除标志（0代表存在 2代表删除）';


--
-- TOC entry 3610 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.login_ip; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.login_ip IS '最后登录IP';


--
-- TOC entry 3611 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.login_date; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.login_date IS '最后登录时间';


--
-- TOC entry 3612 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.create_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.create_by IS '创建者';


--
-- TOC entry 3613 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.create_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.create_time IS '创建时间';


--
-- TOC entry 3614 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.update_by; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.update_by IS '更新者';


--
-- TOC entry 3615 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.update_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.update_time IS '更新时间';


--
-- TOC entry 3616 (class 0 OID 0)
-- Dependencies: 223
-- Name: COLUMN sys_user.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user.remark IS '备注';


--
-- TOC entry 224 (class 1259 OID 24646)
-- Name: sys_user_post; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_user_post (
    user_id bigint NOT NULL,
    post_id bigint NOT NULL
);


ALTER TABLE public.sys_user_post OWNER TO postgres;

--
-- TOC entry 3617 (class 0 OID 0)
-- Dependencies: 224
-- Name: TABLE sys_user_post; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_user_post IS '用户与岗位关联表';


--
-- TOC entry 3618 (class 0 OID 0)
-- Dependencies: 224
-- Name: COLUMN sys_user_post.user_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user_post.user_id IS '用户ID';


--
-- TOC entry 3619 (class 0 OID 0)
-- Dependencies: 224
-- Name: COLUMN sys_user_post.post_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user_post.post_id IS '岗位ID';


--
-- TOC entry 225 (class 1259 OID 24649)
-- Name: sys_user_role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_user_role (
    user_id bigint NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.sys_user_role OWNER TO postgres;

--
-- TOC entry 3620 (class 0 OID 0)
-- Dependencies: 225
-- Name: TABLE sys_user_role; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_user_role IS '用户和角色关联表';


--
-- TOC entry 3621 (class 0 OID 0)
-- Dependencies: 225
-- Name: COLUMN sys_user_role.user_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user_role.user_id IS '用户ID';


--
-- TOC entry 3622 (class 0 OID 0)
-- Dependencies: 225
-- Name: COLUMN sys_user_role.role_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user_role.role_id IS '角色ID';


--
-- TOC entry 3409 (class 0 OID 24577)
-- Dependencies: 209
-- Data for Name: sys_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_config (config_id, config_name, config_key, config_value, config_type, create_by, create_time, update_by, update_time, remark) FROM stdin;
1	主框架页-默认皮肤样式名称	sys.index.skinName	skin-blue	Y	admin	2024-06-01 14:43:36		\N	蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow
2	用户管理-账号初始密码	sys.user.initPassword	123456	Y	admin	2024-06-01 14:43:36		\N	初始化密码 123456
3	主框架页-侧边栏主题	sys.index.sideTheme	theme-dark	Y	admin	2024-06-01 14:43:36		\N	深色主题theme-dark，浅色主题theme-light
4	账号自助-验证码开关	sys.account.captchaEnabled	true	Y	admin	2024-06-01 14:43:36		\N	是否开启验证码功能（true开启，false关闭）
5	账号自助-是否开启用户注册功能	sys.account.registerUser	false	Y	admin	2024-06-01 14:43:36		\N	是否开启注册用户功能（true开启，false关闭）
6	用户登录-黑名单列表	sys.login.blackIPList		Y	admin	2024-06-01 14:43:36		\N	设置登录IP黑名单限制，多个匹配项以;分隔，支持匹配（*通配、网段）
\.


--
-- TOC entry 3410 (class 0 OID 24582)
-- Dependencies: 210
-- Data for Name: sys_dept; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_dept (dept_id, parent_id, ancestors, dept_name, order_num, leader, phone, email, status, del_flag, create_by, create_time, update_by, update_time) FROM stdin;
100	0	0	若依科技	0	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
101	100	0,100	深圳总公司	1	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
102	100	0,100	长沙分公司	2	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
103	101	0,100,101	研发部门	1	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
104	101	0,100,101	市场部门	2	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
105	101	0,100,101	测试部门	3	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
106	101	0,100,101	财务部门	4	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
107	101	0,100,101	运维部门	5	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
108	102	0,100,102	市场部门	1	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
109	102	0,100,102	财务部门	2	若依	15888888888	ry@qq.com	0	0	admin	2024-06-01 14:43:35		\N
\.


--
-- TOC entry 3411 (class 0 OID 24585)
-- Dependencies: 211
-- Data for Name: sys_dict_data; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_dict_data (dict_code, dict_sort, dict_label, dict_value, dict_type, css_class, list_class, is_default, status, create_by, create_time, update_by, update_time, remark) FROM stdin;
1	1	男	0	sys_user_sex			Y	0	admin	2024-06-01 14:43:36		\N	性别男
2	2	女	1	sys_user_sex			N	0	admin	2024-06-01 14:43:36		\N	性别女
3	3	未知	2	sys_user_sex			N	0	admin	2024-06-01 14:43:36		\N	性别未知
4	1	显示	0	sys_show_hide		primary	Y	0	admin	2024-06-01 14:43:36		\N	显示菜单
5	2	隐藏	1	sys_show_hide		danger	N	0	admin	2024-06-01 14:43:36		\N	隐藏菜单
6	1	正常	0	sys_normal_disable		primary	Y	0	admin	2024-06-01 14:43:36		\N	正常状态
7	2	停用	1	sys_normal_disable		danger	N	0	admin	2024-06-01 14:43:36		\N	停用状态
8	1	正常	0	sys_job_status		primary	Y	0	admin	2024-06-01 14:43:36		\N	正常状态
9	2	暂停	1	sys_job_status		danger	N	0	admin	2024-06-01 14:43:36		\N	停用状态
10	1	默认	DEFAULT	sys_job_group			Y	0	admin	2024-06-01 14:43:36		\N	默认分组
11	2	系统	SYSTEM	sys_job_group			N	0	admin	2024-06-01 14:43:36		\N	系统分组
12	1	是	Y	sys_yes_no		primary	Y	0	admin	2024-06-01 14:43:36		\N	系统默认是
13	2	否	N	sys_yes_no		danger	N	0	admin	2024-06-01 14:43:36		\N	系统默认否
14	1	通知	1	sys_notice_type		warning	Y	0	admin	2024-06-01 14:43:36		\N	通知
15	2	公告	2	sys_notice_type		success	N	0	admin	2024-06-01 14:43:36		\N	公告
16	1	正常	0	sys_notice_status		primary	Y	0	admin	2024-06-01 14:43:36		\N	正常状态
17	2	关闭	1	sys_notice_status		danger	N	0	admin	2024-06-01 14:43:36		\N	关闭状态
18	99	其他	0	sys_oper_type		info	N	0	admin	2024-06-01 14:43:36		\N	其他操作
19	1	新增	1	sys_oper_type		info	N	0	admin	2024-06-01 14:43:36		\N	新增操作
20	2	修改	2	sys_oper_type		info	N	0	admin	2024-06-01 14:43:36		\N	修改操作
21	3	删除	3	sys_oper_type		danger	N	0	admin	2024-06-01 14:43:36		\N	删除操作
22	4	授权	4	sys_oper_type		primary	N	0	admin	2024-06-01 14:43:36		\N	授权操作
23	5	导出	5	sys_oper_type		warning	N	0	admin	2024-06-01 14:43:36		\N	导出操作
24	6	导入	6	sys_oper_type		warning	N	0	admin	2024-06-01 14:43:36		\N	导入操作
25	7	强退	7	sys_oper_type		danger	N	0	admin	2024-06-01 14:43:36		\N	强退操作
26	8	生成代码	8	sys_oper_type		warning	N	0	admin	2024-06-01 14:43:36		\N	生成操作
27	9	清空数据	9	sys_oper_type		danger	N	0	admin	2024-06-01 14:43:36		\N	清空操作
28	1	成功	0	sys_common_status		primary	N	0	admin	2024-06-01 14:43:36		\N	正常状态
29	2	失败	1	sys_common_status		danger	N	0	admin	2024-06-01 14:43:36		\N	停用状态
\.


--
-- TOC entry 3412 (class 0 OID 24590)
-- Dependencies: 212
-- Data for Name: sys_dict_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_dict_type (dict_id, dict_name, dict_type, status, create_by, create_time, update_by, update_time, remark) FROM stdin;
1	用户性别	sys_user_sex	0	admin	2024-06-01 14:43:36		\N	用户性别列表
2	菜单状态	sys_show_hide	0	admin	2024-06-01 14:43:36		\N	菜单状态列表
3	系统开关	sys_normal_disable	0	admin	2024-06-01 14:43:36		\N	系统开关列表
4	任务状态	sys_job_status	0	admin	2024-06-01 14:43:36		\N	任务状态列表
5	任务分组	sys_job_group	0	admin	2024-06-01 14:43:36		\N	任务分组列表
6	系统是否	sys_yes_no	0	admin	2024-06-01 14:43:36		\N	系统是否列表
7	通知类型	sys_notice_type	0	admin	2024-06-01 14:43:36		\N	通知类型列表
8	通知状态	sys_notice_status	0	admin	2024-06-01 14:43:36		\N	通知状态列表
9	操作类型	sys_oper_type	0	admin	2024-06-01 14:43:36		\N	操作类型列表
10	系统状态	sys_common_status	0	admin	2024-06-01 14:43:36		\N	登录状态列表
\.


--
-- TOC entry 3413 (class 0 OID 24595)
-- Dependencies: 213
-- Data for Name: sys_job; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_job (job_id, job_name, job_group, invoke_target, cron_expression, misfire_policy, concurrent, status, create_by, create_time, update_by, update_time, remark) FROM stdin;
1	系统默认（无参）	DEFAULT	ryTask.ryNoParams	0/10 * * * * ?	3	1	1	admin	2024-06-01 14:43:36		\N	
2	系统默认（有参）	DEFAULT	ryTask.ryParams('ry')	0/15 * * * * ?	3	1	1	admin	2024-06-01 14:43:36		\N	
3	系统默认（多参）	DEFAULT	ryTask.ryMultipleParams('ry', true, 2000L, 316.50D, 100)	0/20 * * * * ?	3	1	1	admin	2024-06-01 14:43:36		\N	
\.


--
-- TOC entry 3414 (class 0 OID 24600)
-- Dependencies: 214
-- Data for Name: sys_job_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_job_log (job_log_id, job_name, job_group, invoke_target, job_message, status, exception_info, create_time) FROM stdin;
\.


--
-- TOC entry 3415 (class 0 OID 24605)
-- Dependencies: 215
-- Data for Name: sys_logininfor; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_logininfor (info_id, user_name, ipaddr, login_location, browser, os, status, msg, login_time) FROM stdin;
\.


--
-- TOC entry 3416 (class 0 OID 24610)
-- Dependencies: 216
-- Data for Name: sys_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_menu (menu_id, menu_name, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, create_by, create_time, update_by, update_time, remark) FROM stdin;
1	系统管理	0	1	system	\N		1	0	M	0	0		system	admin	2024-06-01 14:43:36		\N	系统管理目录
2	系统监控	0	2	monitor	\N		1	0	M	0	0		monitor	admin	2024-06-01 14:43:36		\N	系统监控目录
3	系统工具	0	3	tool	\N		1	0	M	0	0		tool	admin	2024-06-01 14:43:36		\N	系统工具目录
4	若依官网	0	4	http://ruoyi.vip	\N		0	0	M	0	0		guide	admin	2024-06-01 14:43:36		\N	若依官网地址
100	用户管理	1	1	user	system/user/index		1	0	C	0	0	system:user:list	user	admin	2024-06-01 14:43:36		\N	用户管理菜单
101	角色管理	1	2	role	system/role/index		1	0	C	0	0	system:role:list	peoples	admin	2024-06-01 14:43:36		\N	角色管理菜单
102	菜单管理	1	3	menu	system/menu/index		1	0	C	0	0	system:menu:list	tree-table	admin	2024-06-01 14:43:36		\N	菜单管理菜单
103	部门管理	1	4	dept	system/dept/index		1	0	C	0	0	system:dept:list	tree	admin	2024-06-01 14:43:36		\N	部门管理菜单
104	岗位管理	1	5	post	system/post/index		1	0	C	0	0	system:post:list	post	admin	2024-06-01 14:43:36		\N	岗位管理菜单
105	字典管理	1	6	dict	system/dict/index		1	0	C	0	0	system:dict:list	dict	admin	2024-06-01 14:43:36		\N	字典管理菜单
106	参数设置	1	7	config	system/config/index		1	0	C	0	0	system:config:list	edit	admin	2024-06-01 14:43:36		\N	参数设置菜单
107	通知公告	1	8	notice	system/notice/index		1	0	C	0	0	system:notice:list	message	admin	2024-06-01 14:43:36		\N	通知公告菜单
108	日志管理	1	9	log			1	0	M	0	0		log	admin	2024-06-01 14:43:36		\N	日志管理菜单
109	在线用户	2	1	online	monitor/online/index		1	0	C	0	0	monitor:online:list	online	admin	2024-06-01 14:43:36		\N	在线用户菜单
110	定时任务	2	2	job	monitor/job/index		1	0	C	0	0	monitor:job:list	job	admin	2024-06-01 14:43:36		\N	定时任务菜单
111	数据监控	2	3	druid	monitor/druid/index		1	0	C	0	0	monitor:druid:list	druid	admin	2024-06-01 14:43:36		\N	数据监控菜单
112	服务监控	2	4	server	monitor/server/index		1	0	C	0	0	monitor:server:list	server	admin	2024-06-01 14:43:36		\N	服务监控菜单
113	缓存监控	2	5	cache	monitor/cache/index		1	0	C	0	0	monitor:cache:list	redis	admin	2024-06-01 14:43:36		\N	缓存监控菜单
114	缓存列表	2	6	cacheList	monitor/cache/list		1	0	C	0	0	monitor:cache:list	redis-list	admin	2024-06-01 14:43:36		\N	缓存列表菜单
115	表单构建	3	1	build	tool/build/index		1	0	C	0	0	tool:build:list	build	admin	2024-06-01 14:43:36		\N	表单构建菜单
116	代码生成	3	2	gen	tool/gen/index		1	0	C	0	0	tool:gen:list	code	admin	2024-06-01 14:43:36		\N	代码生成菜单
117	系统接口	3	3	swagger	tool/swagger/index		1	0	C	0	0	tool:swagger:list	swagger	admin	2024-06-01 14:43:36		\N	系统接口菜单
500	操作日志	108	1	operlog	monitor/operlog/index		1	0	C	0	0	monitor:operlog:list	form	admin	2024-06-01 14:43:36		\N	操作日志菜单
501	登录日志	108	2	logininfor	monitor/logininfor/index		1	0	C	0	0	monitor:logininfor:list	logininfor	admin	2024-06-01 14:43:36		\N	登录日志菜单
1000	用户查询	100	1				1	0	F	0	0	system:user:query	#	admin	2024-06-01 14:43:36		\N	
1001	用户新增	100	2				1	0	F	0	0	system:user:add	#	admin	2024-06-01 14:43:36		\N	
1002	用户修改	100	3				1	0	F	0	0	system:user:edit	#	admin	2024-06-01 14:43:36		\N	
1003	用户删除	100	4				1	0	F	0	0	system:user:remove	#	admin	2024-06-01 14:43:36		\N	
1004	用户导出	100	5				1	0	F	0	0	system:user:export	#	admin	2024-06-01 14:43:36		\N	
1005	用户导入	100	6				1	0	F	0	0	system:user:import	#	admin	2024-06-01 14:43:36		\N	
1006	重置密码	100	7				1	0	F	0	0	system:user:resetPwd	#	admin	2024-06-01 14:43:36		\N	
1007	角色查询	101	1				1	0	F	0	0	system:role:query	#	admin	2024-06-01 14:43:36		\N	
1008	角色新增	101	2				1	0	F	0	0	system:role:add	#	admin	2024-06-01 14:43:36		\N	
1009	角色修改	101	3				1	0	F	0	0	system:role:edit	#	admin	2024-06-01 14:43:36		\N	
1010	角色删除	101	4				1	0	F	0	0	system:role:remove	#	admin	2024-06-01 14:43:36		\N	
1011	角色导出	101	5				1	0	F	0	0	system:role:export	#	admin	2024-06-01 14:43:36		\N	
1012	菜单查询	102	1				1	0	F	0	0	system:menu:query	#	admin	2024-06-01 14:43:36		\N	
1013	菜单新增	102	2				1	0	F	0	0	system:menu:add	#	admin	2024-06-01 14:43:36		\N	
1014	菜单修改	102	3				1	0	F	0	0	system:menu:edit	#	admin	2024-06-01 14:43:36		\N	
1015	菜单删除	102	4				1	0	F	0	0	system:menu:remove	#	admin	2024-06-01 14:43:36		\N	
1016	部门查询	103	1				1	0	F	0	0	system:dept:query	#	admin	2024-06-01 14:43:36		\N	
1017	部门新增	103	2				1	0	F	0	0	system:dept:add	#	admin	2024-06-01 14:43:36		\N	
1018	部门修改	103	3				1	0	F	0	0	system:dept:edit	#	admin	2024-06-01 14:43:36		\N	
1019	部门删除	103	4				1	0	F	0	0	system:dept:remove	#	admin	2024-06-01 14:43:36		\N	
1020	岗位查询	104	1				1	0	F	0	0	system:post:query	#	admin	2024-06-01 14:43:36		\N	
1021	岗位新增	104	2				1	0	F	0	0	system:post:add	#	admin	2024-06-01 14:43:36		\N	
1022	岗位修改	104	3				1	0	F	0	0	system:post:edit	#	admin	2024-06-01 14:43:36		\N	
1023	岗位删除	104	4				1	0	F	0	0	system:post:remove	#	admin	2024-06-01 14:43:36		\N	
1024	岗位导出	104	5				1	0	F	0	0	system:post:export	#	admin	2024-06-01 14:43:36		\N	
1025	字典查询	105	1	#			1	0	F	0	0	system:dict:query	#	admin	2024-06-01 14:43:36		\N	
1026	字典新增	105	2	#			1	0	F	0	0	system:dict:add	#	admin	2024-06-01 14:43:36		\N	
1027	字典修改	105	3	#			1	0	F	0	0	system:dict:edit	#	admin	2024-06-01 14:43:36		\N	
1028	字典删除	105	4	#			1	0	F	0	0	system:dict:remove	#	admin	2024-06-01 14:43:36		\N	
1029	字典导出	105	5	#			1	0	F	0	0	system:dict:export	#	admin	2024-06-01 14:43:36		\N	
1030	参数查询	106	1	#			1	0	F	0	0	system:config:query	#	admin	2024-06-01 14:43:36		\N	
1031	参数新增	106	2	#			1	0	F	0	0	system:config:add	#	admin	2024-06-01 14:43:36		\N	
1032	参数修改	106	3	#			1	0	F	0	0	system:config:edit	#	admin	2024-06-01 14:43:36		\N	
1033	参数删除	106	4	#			1	0	F	0	0	system:config:remove	#	admin	2024-06-01 14:43:36		\N	
1034	参数导出	106	5	#			1	0	F	0	0	system:config:export	#	admin	2024-06-01 14:43:36		\N	
1035	公告查询	107	1	#			1	0	F	0	0	system:notice:query	#	admin	2024-06-01 14:43:36		\N	
1036	公告新增	107	2	#			1	0	F	0	0	system:notice:add	#	admin	2024-06-01 14:43:36		\N	
1037	公告修改	107	3	#			1	0	F	0	0	system:notice:edit	#	admin	2024-06-01 14:43:36		\N	
1038	公告删除	107	4	#			1	0	F	0	0	system:notice:remove	#	admin	2024-06-01 14:43:36		\N	
1039	操作查询	500	1	#			1	0	F	0	0	monitor:operlog:query	#	admin	2024-06-01 14:43:36		\N	
1040	操作删除	500	2	#			1	0	F	0	0	monitor:operlog:remove	#	admin	2024-06-01 14:43:36		\N	
1041	日志导出	500	3	#			1	0	F	0	0	monitor:operlog:export	#	admin	2024-06-01 14:43:36		\N	
1042	登录查询	501	1	#			1	0	F	0	0	monitor:logininfor:query	#	admin	2024-06-01 14:43:36		\N	
1043	登录删除	501	2	#			1	0	F	0	0	monitor:logininfor:remove	#	admin	2024-06-01 14:43:36		\N	
1044	日志导出	501	3	#			1	0	F	0	0	monitor:logininfor:export	#	admin	2024-06-01 14:43:36		\N	
1045	账户解锁	501	4	#			1	0	F	0	0	monitor:logininfor:unlock	#	admin	2024-06-01 14:43:36		\N	
1046	在线查询	109	1	#			1	0	F	0	0	monitor:online:query	#	admin	2024-06-01 14:43:36		\N	
1047	批量强退	109	2	#			1	0	F	0	0	monitor:online:batchLogout	#	admin	2024-06-01 14:43:36		\N	
1048	单条强退	109	3	#			1	0	F	0	0	monitor:online:forceLogout	#	admin	2024-06-01 14:43:36		\N	
1049	任务查询	110	1	#			1	0	F	0	0	monitor:job:query	#	admin	2024-06-01 14:43:36		\N	
1050	任务新增	110	2	#			1	0	F	0	0	monitor:job:add	#	admin	2024-06-01 14:43:36		\N	
1051	任务修改	110	3	#			1	0	F	0	0	monitor:job:edit	#	admin	2024-06-01 14:43:36		\N	
1052	任务删除	110	4	#			1	0	F	0	0	monitor:job:remove	#	admin	2024-06-01 14:43:36		\N	
1053	状态修改	110	5	#			1	0	F	0	0	monitor:job:changeStatus	#	admin	2024-06-01 14:43:36		\N	
1054	任务导出	110	6	#			1	0	F	0	0	monitor:job:export	#	admin	2024-06-01 14:43:36		\N	
1055	生成查询	116	1	#			1	0	F	0	0	tool:gen:query	#	admin	2024-06-01 14:43:36		\N	
1056	生成修改	116	2	#			1	0	F	0	0	tool:gen:edit	#	admin	2024-06-01 14:43:36		\N	
1057	生成删除	116	3	#			1	0	F	0	0	tool:gen:remove	#	admin	2024-06-01 14:43:36		\N	
1058	导入代码	116	4	#			1	0	F	0	0	tool:gen:import	#	admin	2024-06-01 14:43:36		\N	
1059	预览代码	116	5	#			1	0	F	0	0	tool:gen:preview	#	admin	2024-06-01 14:43:36		\N	
1060	生成代码	116	6	#			1	0	F	0	0	tool:gen:code	#	admin	2024-06-01 14:43:36		\N	
\.


--
-- TOC entry 3417 (class 0 OID 24615)
-- Dependencies: 217
-- Data for Name: sys_notice; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_notice (notice_id, notice_title, notice_type, notice_content, status, create_by, create_time, update_by, update_time, remark) FROM stdin;
1	温馨提醒：2018-07-01 若依新版本发布啦	2	\\xe696b0e78988e69cace58685e5aeb9	0	admin	2024-06-01 14:43:36		\N	管理员
2	维护通知：2018-07-01 若依系统凌晨维护	1	\\xe7bbb4e68aa4e58685e5aeb9	0	admin	2024-06-01 14:43:36		\N	管理员
\.


--
-- TOC entry 3418 (class 0 OID 24620)
-- Dependencies: 218
-- Data for Name: sys_oper_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_oper_log (oper_id, title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time, cost_time) FROM stdin;
\.


--
-- TOC entry 3419 (class 0 OID 24625)
-- Dependencies: 219
-- Data for Name: sys_post; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_post (post_id, post_code, post_name, post_sort, status, create_by, create_time, update_by, update_time, remark) FROM stdin;
1	ceo	董事长	1	0	admin	2024-06-01 14:43:36		\N	
2	se	项目经理	2	0	admin	2024-06-01 14:43:36		\N	
3	hr	人力资源	3	0	admin	2024-06-01 14:43:36		\N	
4	user	普通员工	4	0	admin	2024-06-01 14:43:36		\N	
\.


--
-- TOC entry 3420 (class 0 OID 24630)
-- Dependencies: 220
-- Data for Name: sys_role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_role (role_id, role_name, role_key, role_sort, data_scope, menu_check_strictly, dept_check_strictly, status, del_flag, create_by, create_time, update_by, update_time, remark) FROM stdin;
1	超级管理员	admin	1	1	1	1	0	0	admin	2024-06-01 14:43:36		\N	超级管理员
2	普通角色	common	2	2	1	1	0	0	admin	2024-06-01 14:43:36		\N	普通角色
\.


--
-- TOC entry 3421 (class 0 OID 24635)
-- Dependencies: 221
-- Data for Name: sys_role_dept; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_role_dept (role_id, dept_id) FROM stdin;
2	100
2	101
2	105
\.


--
-- TOC entry 3422 (class 0 OID 24638)
-- Dependencies: 222
-- Data for Name: sys_role_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_role_menu (role_id, menu_id) FROM stdin;
2	1
2	2
2	3
2	4
2	100
2	101
2	102
2	103
2	104
2	105
2	106
2	107
2	108
2	109
2	110
2	111
2	112
2	113
2	114
2	115
2	116
2	117
2	500
2	501
2	1000
2	1001
2	1002
2	1003
2	1004
2	1005
2	1006
2	1007
2	1008
2	1009
2	1010
2	1011
2	1012
2	1013
2	1014
2	1015
2	1016
2	1017
2	1018
2	1019
2	1020
2	1021
2	1022
2	1023
2	1024
2	1025
2	1026
2	1027
2	1028
2	1029
2	1030
2	1031
2	1032
2	1033
2	1034
2	1035
2	1036
2	1037
2	1038
2	1039
2	1040
2	1041
2	1042
2	1043
2	1044
2	1045
2	1046
2	1047
2	1048
2	1049
2	1050
2	1051
2	1052
2	1053
2	1054
2	1055
2	1056
2	1057
2	1058
2	1059
2	1060
\.


--
-- TOC entry 3423 (class 0 OID 24641)
-- Dependencies: 223
-- Data for Name: sys_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_user (user_id, dept_id, user_name, nick_name, user_type, email, phonenumber, sex, avatar, password, status, del_flag, login_ip, login_date, create_by, create_time, update_by, update_time, remark) FROM stdin;
1	103	admin	若依	00	ry@163.com	15888888888	1		$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2	0	0	127.0.0.1	2024-06-01 14:43:35	admin	2024-06-01 14:43:35		\N	管理员
2	105	ry	若依	00	ry@qq.com	15666666666	1		$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2	0	0	127.0.0.1	2024-06-01 14:43:35	admin	2024-06-01 14:43:35		\N	测试员
\.


--
-- TOC entry 3424 (class 0 OID 24646)
-- Dependencies: 224
-- Data for Name: sys_user_post; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_user_post (user_id, post_id) FROM stdin;
1	1
2	2
\.


--
-- TOC entry 3425 (class 0 OID 24649)
-- Dependencies: 225
-- Data for Name: sys_user_role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_user_role (user_id, role_id) FROM stdin;
1	1
1	2
2	2
\.


--
-- TOC entry 3231 (class 2606 OID 24653)
-- Name: sys_config sys_config_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_config
    ADD CONSTRAINT sys_config_pkey PRIMARY KEY (config_id);


--
-- TOC entry 3233 (class 2606 OID 24655)
-- Name: sys_dept sys_dept_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dept
    ADD CONSTRAINT sys_dept_pkey PRIMARY KEY (dept_id);


--
-- TOC entry 3235 (class 2606 OID 24657)
-- Name: sys_dict_data sys_dict_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dict_data
    ADD CONSTRAINT sys_dict_data_pkey PRIMARY KEY (dict_code);


--
-- TOC entry 3238 (class 2606 OID 24660)
-- Name: sys_dict_type sys_dict_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dict_type
    ADD CONSTRAINT sys_dict_type_pkey PRIMARY KEY (dict_id);


--
-- TOC entry 3242 (class 2606 OID 24664)
-- Name: sys_job_log sys_job_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_job_log
    ADD CONSTRAINT sys_job_log_pkey PRIMARY KEY (job_log_id);


--
-- TOC entry 3240 (class 2606 OID 24662)
-- Name: sys_job sys_job_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_job
    ADD CONSTRAINT sys_job_pkey PRIMARY KEY (job_id, job_name, job_group);


--
-- TOC entry 3246 (class 2606 OID 24668)
-- Name: sys_logininfor sys_logininfor_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_logininfor
    ADD CONSTRAINT sys_logininfor_pkey PRIMARY KEY (info_id);


--
-- TOC entry 3248 (class 2606 OID 24670)
-- Name: sys_menu sys_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_menu
    ADD CONSTRAINT sys_menu_pkey PRIMARY KEY (menu_id);


--
-- TOC entry 3250 (class 2606 OID 24672)
-- Name: sys_notice sys_notice_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_notice
    ADD CONSTRAINT sys_notice_pkey PRIMARY KEY (notice_id);


--
-- TOC entry 3255 (class 2606 OID 24677)
-- Name: sys_oper_log sys_oper_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_oper_log
    ADD CONSTRAINT sys_oper_log_pkey PRIMARY KEY (oper_id);


--
-- TOC entry 3257 (class 2606 OID 24679)
-- Name: sys_post sys_post_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_post
    ADD CONSTRAINT sys_post_pkey PRIMARY KEY (post_id);


--
-- TOC entry 3261 (class 2606 OID 24683)
-- Name: sys_role_dept sys_role_dept_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_role_dept
    ADD CONSTRAINT sys_role_dept_pkey PRIMARY KEY (role_id, dept_id);


--
-- TOC entry 3263 (class 2606 OID 24685)
-- Name: sys_role_menu sys_role_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_role_menu
    ADD CONSTRAINT sys_role_menu_pkey PRIMARY KEY (role_id, menu_id);


--
-- TOC entry 3259 (class 2606 OID 24681)
-- Name: sys_role sys_role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_role
    ADD CONSTRAINT sys_role_pkey PRIMARY KEY (role_id);


--
-- TOC entry 3265 (class 2606 OID 24687)
-- Name: sys_user sys_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_user
    ADD CONSTRAINT sys_user_pkey PRIMARY KEY (user_id);


--
-- TOC entry 3267 (class 2606 OID 24689)
-- Name: sys_user_post sys_user_post_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_user_post
    ADD CONSTRAINT sys_user_post_pkey PRIMARY KEY (user_id, post_id);


--
-- TOC entry 3269 (class 2606 OID 24691)
-- Name: sys_user_role sys_user_role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_user_role
    ADD CONSTRAINT sys_user_role_pkey PRIMARY KEY (user_id, role_id);


--
-- TOC entry 3236 (class 1259 OID 24658)
-- Name: dict_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX dict_type ON public.sys_dict_type USING btree (dict_type);


--
-- TOC entry 3243 (class 1259 OID 24666)
-- Name: idx_sys_logininfor_lt; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_logininfor_lt ON public.sys_logininfor USING btree (login_time);


--
-- TOC entry 3244 (class 1259 OID 24665)
-- Name: idx_sys_logininfor_s; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_logininfor_s ON public.sys_logininfor USING btree (status);


--
-- TOC entry 3251 (class 1259 OID 24673)
-- Name: idx_sys_oper_log_bt; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_oper_log_bt ON public.sys_oper_log USING btree (business_type);


--
-- TOC entry 3252 (class 1259 OID 24675)
-- Name: idx_sys_oper_log_ot; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_oper_log_ot ON public.sys_oper_log USING btree (oper_time);


--
-- TOC entry 3253 (class 1259 OID 24674)
-- Name: idx_sys_oper_log_s; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_oper_log_s ON public.sys_oper_log USING btree (status);


--
-- TOC entry 3431 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2024-06-15 16:36:51

--
-- PostgreSQL database dump complete
--

-- Completed on 2024-06-15 16:36:51

--
-- PostgreSQL database cluster dump complete
--

