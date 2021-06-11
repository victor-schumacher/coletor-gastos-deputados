-- CRIAÇÃO ESQUEMA DEPUTADOS

create schema deputados;

-- CRIAÇÃO TABELA GASTOS

create table deputados.gastos
(
    id                UUID,
    datemissao        timestamp,
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

-- BigNumber valor gasto no dia anterior
select
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -1


-- Ranking de deputados que mais gastam
select
	txnomeparlamentar as "Deputado",
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -10000
group by 1
order by 2 desc


-- Visão agrupada por período de deputados
select
	date_trunc('day', datemissao) as "Data",
	txnomeparlamentar as "Deputado",
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -10000
group by 1, 2
order by 1, 2, 3


-- Ranking partidos que mais gastam
select
	sgpartido as "Partido",
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -10000
group by 1
order by 2 desc


-- Visão agrupada por período de partidos
select
	date_trunc('day', datemissao) as "Data",
	sgpartido as "Partido",
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -10000
group by 1, 2
order by 1, 2, 3


-- Ranking tipos de gastos
select
	txtdescricao as "Nome do gasto",
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -10000
group by 1
order by 2 desc


-- Visão agrupada por período de tipos de gastos
select
	date_trunc('day', datemissao) as "Data",
	txtdescricao as "Nome do gasto",
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -10000
group by 1, 2
order by 1, 2, 3


-- Ranking fornecedores
select
	txtfornecedor as "Nome do fornecedor",
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -10000
group by 1
order by 2 desc


-- Visão agrupada por período de fornecedores
select
	date_trunc('day', datemissao) as "Data",
	txtfornecedor as "Nome do fornecedor",
	sum(vlrliquido) as "Valor"
from deputados.gastos
where datemissao >= current_date -10000
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


