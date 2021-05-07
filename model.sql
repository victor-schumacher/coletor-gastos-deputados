create schema deputados;

create table deputados.gastos (
	id bigserial,
	datemissao timestamp,
	nulegislatura text,
	sgpartido  text,
	txnomeparlamentar text,
	txtcnpjcpf text,
	txtdescricao text,
	txtfornecedor text,
	vlrliquido float
)

create index idx_datemissao ON deputados.gastos (datemissao);
create index idx_sgpartido ON deputados.gastos (sgpartido);
create index idx_txnomeparlamentar ON deputados.gastos (txnomeparlamentar);
