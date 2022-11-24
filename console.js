const createUseCase = require('./.console/createUseCase');
const createRepository = require('./.console/createRepository');
const [, , action, ...args] = process.argv

switch (action) {
  case 'usecase':
  case 'u':
    createUseCase(args);
    break;
  case 'repository':
  case 'r':
    createRepository(args);
    break;
}