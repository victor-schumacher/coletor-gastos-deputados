-- CRIAÇÃO ESQUEMA DEPUTADOS

create schema deputados;

-- CRIAÇÃO TABELA GASTOS

create table gastos (
    id                UUID,
    datemissao        text,
    nulegislatura     text,
    sgpartido         text,
    txnomeparlamentar text,
    txtcnpjcpf        text,
    txtdescricao      text,
    txtfornecedor     text,
    vlrliquido        float
);

create index idx_datemissao ON deputados.gastos (datemissao);
create index idx_sgpartido ON deputados.gastos (sgpartido);
create index idx_txnomeparlamentar ON deputados.gastos (txnomeparlamentar);


-- VISÕES METABASE
-- Escopo: https://trello.com/c/APekaIaL/19-levantar-escopo-vis%C3%B5es-do-dashboard
-- Moqup: https://trello.com/c/mqJhXbbQ/18-construir-moqup-vis%C3%B5es-dashboards

-- Soma do valor gasto no período
select
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
    
    
-- Soma do valor gasto no período por mês
select
	date_trunc('month', datemissao::timestamp) as "Data",
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1
order by 1


-- Top 10 deputados que mais gastam
select
	txnomeparlamentar as "Deputado",
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1
order by 2 desc
limit 10

-- Soma do valor gasto no período por deputados por mês
with base_filter as(
select
	txnomeparlamentar,
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1
order by 2 desc
limit 10
)
select
	date_trunc('month', datemissao::timestamp) as "Data",
	concat(gastos.txnomeparlamentar, ' - ', sgpartido) as "Deputado",
	sum(vlrliquido) as "Valor"
from gastos
join base_filter
    on base_filter.txnomeparlamentar = gastos.txnomeparlamentar
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1, 2
order by 1, 2, 3


-- Top 10 partidos que mais gastam
select
	sgpartido as "Partido",
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1
order by 2 desc
limit 10


-- Soma do valor gasto no período por partidos por mês
with base_filter as(
select
	sgpartido,
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1
order by 2 desc
limit 10
)
select
	date_trunc('month', datemissao::timestamp) as "Data",
	gastos.sgpartido as "Partido",
	sum(vlrliquido) as "Valor"
from gastos
join base_filter
    on base_filter.sgpartido = gastos.sgpartido
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1, 2
order by 1, 2, 3


-- Top 10 tipos de gasto que mais gastam
select
	txtdescricao as "Nome do gasto",
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1
order by 2 desc
limit 10


-- Soma do valor gasto no período por tipos de gasto por mês
with base_filter as(
select
	txtdescricao,
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1
order by 2 desc
limit 10
)
select
	date_trunc('month', datemissao::timestamp) as "Data",
	gastos.txtdescricao as "Nome do gasto",
	sum(vlrliquido) as "Valor"
from gastos
join base_filter
    on base_filter.txtdescricao = gastos.txtdescricao
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1, 2
order by 1, 2, 3


-- Top 10 fornecedor que mais receberam
select
	txtfornecedor as "Nome do fornecedor",
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1
order by 2 desc
limit 10


-- Soma do valor recebido no período por fornecedores por mês
with base_filter as(
select
	txtfornecedor,
	sum(vlrliquido) as "Valor"
from gastos
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1
order by 2 desc
limit 10
)
select
	date_trunc('month', datemissao::timestamp) as "Data",
	gastos.txtfornecedor as "Nome do fornecedor",
	sum(vlrliquido) as "Valor"
from gastos
join base_filter
    on base_filter.txtfornecedor = gastos.txtfornecedor
where
    datemissao is not null
    and datemissao <> ''
    and datemissao::timestamp >= {{data_inicio}}
    and datemissao::timestamp <= {{data_fim}}
group by 1, 2
order by 1, 2, 3


-- Dispersão dos maiores gastos X deputados X valores
select
	txtdescricao as "Nome do gasto",
	txnomeparlamentar as "Deputado",
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -10000
group by 1, 2
