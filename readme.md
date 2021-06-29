## [DASHBOARD GASTOS DEPUTADOS BRASIL](https://metabase-catolica-sc.herokuapp.com/public/dashboard/22847b17-113f-4817-82f0-890271bc43ca?data_in%25C3%25ADcio=2018-01-01&data_fim=2020-04-30#theme=night)


### O que é?

Trata-se de um projeto que nasceu na disciplina de fábrica de _software_ do [Centro universitário católica de Santa Catarina](https://www.catolicasc.org.br/) em Joinville/SC e tem como objetivo a extração de dados de gastos de deputados da base disponibilizada pelo [Brasil.io](https://brasil.io/dataset/gastos-deputados/cota_parlamentar/) e a construção de algumas visões com base nos dados extraídos;


### Por quê?

Para trazer uma visão, ainda que mínima, dos gastos que são registrados pelos deputados e pouco divulgados à população brasileira.

### Como é construído?

#### Extração de dados:

Através da integração construída utilizando a linguagem de programação [Golang](https://golang.org), extraímos dados dos arquivos _.csv_ disponibilizados pela [Brasil.io](https://brasil.io/dataset/gastos-deputados/cota_parlamentar/) e fazemos o envio para um banco de dados PostgreSQL.

#### Visões:

Conectamos a aplicação [Metabase](https://www.metabase.com/) ao banco de dados [PostgreSQL](https://www.postgresql.org/) que contém os dados e geramos as visões através da linguagem de consulta [_sql_](https://pt.wikipedia.org/wiki/SQL).


##### Consultas _SQL_ das visões no Metabase:

- [Soma do valor gasto no período](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/598cba090dd35bd49b0cbaea7dab6fd0f5bb8d16/model.sql#L28);
- [Soma do valor gasto no período por mês](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L39);
- [Top 10 deputados que mais gastaram](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L53);
- [Soma do valor gasto no período por deputados por mês](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L67);
- [Top 10 partidos que mais gastaram](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L98);
- [Soma do valor gasto no período por partidos por mês](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L113);
- [Top 10 tipos de gasto que mais gastaram](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L144);
- [Soma do valor gasto no período por tipos de gasto por mês](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L159);
- [Top 10 fornecedores que mais receberam](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L190);
- [Soma do valor recebido no período por fornecedores por mês](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L205);
- [Dispersão dos maiores tipos de gastos X partidos X valores](https://github.com/victor-schumacher/coletor-gastos-deputados/blob/012d4a234a69b98cc27bed6bf680811d8186bb43/model.sql#L236).
