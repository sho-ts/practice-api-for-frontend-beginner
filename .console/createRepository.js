const fs = require('fs');

/** RepositoryのInterfaceと構造体を作成する */
module.exports = ([, tableName]) => {
  const fileName = `${tableName}.go`;
  const upperdTableName = tableName.charAt(0).toUpperCase() + tableName.slice(1);

  const interactorValue = `package repository\n\ntype ${tableName}Repository struct{}\n\nfunc New${upperdTableName}Repository() ${tableName}Repository {\n  return ${tableName}Repository{}\n}`;
  const interfaceValue = `package repository\n\ntype I${upperdTableName}Repository interface {}`;

  fs.writeFileSync(`./repository/${fileName}`, interactorValue);
  fs.writeFileSync(`./domain/repository/${fileName}`, interfaceValue);
}