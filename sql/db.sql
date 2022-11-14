CREATE TABLE IF NOT EXISTS orders
(
    order_uid character varying(36),
    track_number character varying(32),
    entry character varying(10),
    locale character varying(2),
    internal_signature character varying(10),
    customer_id character varying(20),
    delivery_service character varying(50),
    shardkey character varying(20),
    sm_id integer,
    date_created timestamp without time zone,
    oof_shard character varying(20),
    CONSTRAINT orders_pkey PRIMARY KEY (order_uid)
)


CREATE TABLE IF NOT EXISTS delivery
(
    name character varying(20),
    phone character varying(12),
    zip character varying(30),
    city character varying(30),
    address character varying(50),
    region character varying(30),
    email character varying(50),
    order_uid character varying(36),
    CONSTRAINT fk_order FOREIGN KEY (order_uid)
        REFERENCES orders (order_uid)
)

CREATE TABLE IF NOT EXISTS payments
(
    transaction character varying(40),
    request_id character varying(50) ,
    currency character varying(50) ,
    provider character varying(50),
    amount integer,
    payment_dt integer,
    bank character varying(50),
    delivery_cost integer,
    goods_total integer,
    custom_fee integer,
    order_uid character varying(36),
    CONSTRAINT fk_order FOREIGN KEY (order_uid)
        REFERENCES orders (order_uid)
)

CREATE TABLE IF NOT EXISTS public.items
(
    chrt_id integer,
    track_number character varying(30),
    price integer,
    rid character varying(30),
    name character varying(30),
    sale integer,
    size character varying(30),
    total_price integer,
    nm_id integer,
    brand character varying(30),
    status integer,
    order_uid character varying(36),
    CONSTRAINT fk_order FOREIGN KEY (order_uid)
        REFERENCES orders (order_uid)
)