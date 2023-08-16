--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Ubuntu 15.4-1.pgdg20.04+1)
-- Dumped by pg_dump version 15.4 (Ubuntu 15.4-1.pgdg20.04+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: category; Type: TABLE; Schema: public; Owner: vici
--

CREATE TABLE public.category (
    id bigint NOT NULL,
    category_name character varying(50) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone
);


ALTER TABLE public.category OWNER TO vici;

--
-- Name: category_id_seq; Type: SEQUENCE; Schema: public; Owner: vici
--

CREATE SEQUENCE public.category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.category_id_seq OWNER TO vici;

--
-- Name: category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vici
--

ALTER SEQUENCE public.category_id_seq OWNED BY public.category.id;


--
-- Name: product; Type: TABLE; Schema: public; Owner: vici
--

CREATE TABLE public.product (
    id bigint NOT NULL,
    name character varying(50) NOT NULL,
    short_description character varying(120) NOT NULL,
    description character varying(255) NOT NULL,
    price numeric(12,2) NOT NULL,
    quantity bigint NOT NULL,
    created timestamp without time zone NOT NULL,
    updated timestamp without time zone,
    category_id bigint
);


ALTER TABLE public.product OWNER TO vici;

--
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: vici
--

CREATE SEQUENCE public.product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_id_seq OWNER TO vici;

--
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vici
--

ALTER SEQUENCE public.product_id_seq OWNED BY public.product.id;


--
-- Name: category id; Type: DEFAULT; Schema: public; Owner: vici
--

ALTER TABLE ONLY public.category ALTER COLUMN id SET DEFAULT nextval('public.category_id_seq'::regclass);


--
-- Name: product id; Type: DEFAULT; Schema: public; Owner: vici
--

ALTER TABLE ONLY public.product ALTER COLUMN id SET DEFAULT nextval('public.product_id_seq'::regclass);


--
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: vici
--

COPY public.category (id, category_name, created_at, updated_at) FROM stdin;
6	Laptopovi	2023-08-15 13:28:18.72263	2023-08-15 13:30:48.658269
8	Desktop	2023-08-15 14:33:47.311889	0001-01-01 00:00:00
9	Stampaci	2023-08-15 14:33:56.549616	0001-01-01 00:00:00
10	Monitori	2023-08-15 14:34:00.532794	0001-01-01 00:00:00
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: vici
--

COPY public.product (id, name, short_description, description, price, quantity, created, updated, category_id) FROM stdin;
18	Dell XPS 13	Lorem ipsum dolor sit amet	Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam gravida arcu nec sem volutpat porta. Cras vel suscipit lacus, non auctor massa.	1499.99	15	2023-08-15 16:52:56.329429	1970-01-01 00:00:00	6
19	MacBook Pro	Consectetur adipiscing elit	Consectetur adipiscing elit. Nam gravida arcu nec sem volutpat porta. Cras vel suscipit lacus, non auctor massa. Donec id magna nec justo commodo malesuada.	2099.99	27	2023-08-15 16:53:00.092099	1970-01-01 00:00:00	6
20	HP Spectre x360	Sed do eiusmod tempor incididunt	Sed sodales lacus ante, in sagittis leo finibus ac. Donec gravida nisl vel ligula viverra, quis egestas nisl vehicula.	2099.99	27	2023-08-15 16:53:02.959494	1970-01-01 00:00:00	6
21	Gigatron Prime Lider Spark	Lorem ipsum dolor sit amet	Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam gravida arcu nec sem volutpat porta. Cras vel suscipit lacus, non auctor massa.	549.99	14	2023-08-15 16:53:26.350408	1970-01-01 00:00:00	8
22	Gigatron Aurora Pro Base	Consectetur adipiscing elit	Consectetur adipiscing elit. Nam gravida arcu nec sem volutpat porta. Cras vel suscipit lacus, non auctor massa. Donec id magna nec justo commodo malesuada.	379.99	5	2023-08-15 16:53:29.388413	1970-01-01 00:00:00	8
23	Gigatron Prime Lider Paladin	Sed do eiusmod tempor incididunt	Sed sodales lacus ante, in sagittis leo finibus ac. Donec gravida nisl vel ligula viverra, quis egestas nisl vehicula.	619.99	14	2023-08-15 16:53:32.454568	1970-01-01 00:00:00	8
24	PHILIPS Monitor 27 inch V Line	Lorem ipsum dolor sit amet	Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam gravida arcu nec sem volutpat porta. Cras vel suscipit lacus, non auctor massa.	149.99	2	2023-08-15 16:53:35.203629	1970-01-01 00:00:00	10
\.


--
-- Name: category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vici
--

SELECT pg_catalog.setval('public.category_id_seq', 10, true);


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vici
--

SELECT pg_catalog.setval('public.product_id_seq', 24, true);


--
-- Name: category category_pkey; Type: CONSTRAINT; Schema: public; Owner: vici
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: vici
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- Name: product product_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: vici
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.category(id);


--
-- PostgreSQL database dump complete
--

