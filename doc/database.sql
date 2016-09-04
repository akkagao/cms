SET NAMES UTF8;
  create database if not exists cmsadmin default charset utf8 collate utf8_general_ci;
  grant select,update,delete,insert,alter,create,drop on cmsadmin.* to "cmsadmin"@"%" identified by "cmsadmin";
  grant select,update,delete,insert,alter,create,drop on cmsadmin.* to "cmsadmin"@"localhost" identified by "cmsadmin";
USE cmsadmin;