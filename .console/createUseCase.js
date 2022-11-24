const fs = require('fs');

/** UseCase、Interactor、Input, Outputを作成する */
module.exports = ([dirName, useCaseName]) => {
  const fileName = `${useCaseName}.go`;
  const upperduUseCaseName = useCaseName.charAt(0).toUpperCase() + useCaseName.slice(1);

  const useCaseDir = `./domain/useCase/${dirName}`;
  const interactorDir = `./adapter/interactor/${dirName}`;
  const inputDir = `./domain/dto/input/${dirName}`;
  const outputDir = `./domain/dto/output/${dirName}`;

  // それぞれのファイルに書き込む内容を作成する
  const useCaseValue = `package useCase\n\nimport (\n  "note-app/domain/dto/input/${dirName}"\n  "note-app/domain/dto/output/${dirName}"\n)\n\ntype I${upperduUseCaseName}UseCase interface{\n  Handle(input.${upperduUseCaseName}Input) output.${upperduUseCaseName}Output\n}`;
  const interactorValue = `package interactor\n\nimport (\n  "note-app/domain/dto/input/${dirName}"\n  "note-app/domain/dto/output/${dirName}"\n)\n\ntype ${useCaseName}Interactor struct{}\nfunc New${upperduUseCaseName}Interactor() ${useCaseName}Interactor {\n  return ${useCaseName}Interactor{}\n}\n\nfunc (i ${useCaseName}Interactor) Handle(in input.${upperduUseCaseName}Input) output.${upperduUseCaseName}Output {\n  \n}`;
  const inputValue = `package input\n\ntype ${upperduUseCaseName}Input struct{}\n\nfunc New${upperduUseCaseName}Input() ${upperduUseCaseName}Input {\n  return ${upperduUseCaseName}Input{}\n}`;
  const outputValue = `package output\n\ntype ${upperduUseCaseName}Output struct{}\n\nfunc New${upperduUseCaseName}Output() ${upperduUseCaseName}Output {\n  return ${upperduUseCaseName}Output{}\n}`;

  // ディレクトリがない場合、作成
  fs.existsSync(useCaseDir) || fs.mkdirSync(useCaseDir)
  fs.existsSync(interactorDir) || fs.mkdirSync(interactorDir)
  fs.existsSync(inputDir) || fs.mkdirSync(inputDir)
  fs.existsSync(outputDir) || fs.mkdirSync(outputDir)

  fs.writeFileSync(`./${useCaseDir}/${fileName}`, useCaseValue);
  fs.writeFileSync(`./${interactorDir}/${fileName}`, interactorValue);
  fs.writeFileSync(`./${inputDir}/${fileName}`, inputValue);
  fs.writeFileSync(`./${outputDir}/${fileName}`, outputValue);
}