"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const typeorm_1 = require("typeorm");
const models_1 = __importDefault(require("./models"));
const { Test, Test2 } = models_1.default;
const typeOrmConfig = {
    type: "postgres",
    host: "0.0.0.0",
    port: 5432,
    username: "root",
    password: "password",
    database: "typeormTest",
    synchronize: true,
    logging: false,
    entities: [
        Test, Test2
    ]
};
function populateDatabase(connection) {
    return __awaiter(this, void 0, void 0, function* () {
        yield connection.query('DELETE FROM test;');
        yield connection.query('DELETE FROM test2;');
        const test = new Test();
        test.value = "asdf";
        const test2 = new Test();
        test2.value = "zxcv";
        const test3 = new Test2();
        test3.value = 'qwer';
        test.others = [test3];
        test2.others = [test3];
        yield connection.getRepository(Test).save(test);
        yield connection.getRepository(Test).save(test2);
        yield connection.getRepository(Test2).save(test3);
        yield connection.getRepository(Test).delete({ value: "asdf" });
    });
}
(() => __awaiter(void 0, void 0, void 0, function* () {
    const connection = yield typeorm_1.createConnection(typeOrmConfig);
    yield populateDatabase(connection);
    const repo = yield connection.getRepository(Test);
    console.log(...yield repo.find());
    yield connection.close();
}))();
//# sourceMappingURL=index.js.map