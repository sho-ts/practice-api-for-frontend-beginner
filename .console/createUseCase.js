const fs = require('fs');

/** UseCase、Interactor、Input, Outputを作成する */
module.exports = ([, dirName, structName]) => {
  const fileName = `${dirName}/${structName}.go`;
  const upperdStructName = structName.charAt(0).toUpperCase() + structName.slice(1);

  const interactorDir = `./adapter/interactor/${dirName}`;
  const usecaseDir = `./domain/usecase/${dirName}`;
  const inputDir = `./domain/dto/input/${dirName}`;
  const outputDir = `./domain/dto/output/${dirName}`;

  // それぞれのファイルに書き込む内容を作成する
  const usecaseValue = `package usecase\n\ntype I${upperdStructName}UseCase interface{\n  Handle()\n}`;
  const interactorValue = `package interactor\n\ntype ${structName}Interactor struct{}\n\nfunc New${upperdStructName}Interactor() ${structName}Interactor {\n  return ${structName}Interactor{}\n}\n\nfunc (i ${structName}Interactor) Handle() {\n  \n}`;
  const inputValue = `package input\n\ntype ${upperdStructName}Input struct{}\n\nfunc New${upperdStructName}Input() ${upperdStructName}Input {\n  return ${upperdStructName}Input{}\n}`;
  const outputValue = `package output\n\ntype ${upperdStructName}Output struct{}\n\nfunc New${upperdStructName}Output() ${upperdStructName}Output {\n  return ${upperdStructName}Output{}\n}`;

  // ディレクトリがない場合、作成
  fs.existsSync(interactorDir) || fs.mkdirSync(interactorDir)
  fs.existsSync(usecaseDir) || fs.mkdirSync(usecaseDir)
  fs.existsSync(inputDir) || fs.mkdirSync(inputDir)
  fs.existsSync(outputDir) || fs.mkdirSync(outputDir)

  fs.writeFileSync(`./adapter/interactor/${fileName}`, interactorValue);
  fs.writeFileSync(`./domain/usecase/${fileName}`, usecaseValue);
  fs.writeFileSync(`./domain/dto/input/${fileName}`, inputValue);
  fs.writeFileSync(`./domain/dto/output/${fileName}`, outputValue);
}