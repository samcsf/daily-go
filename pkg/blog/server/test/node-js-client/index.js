"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
exports.__esModule = true;
var path = require("path");
var grpc = require("grpc");
var protoLoader = require("@grpc/proto-loader");
var PROTO_PATH = path.join(__dirname, '../../grpc/post/post/post.proto');
var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});
var def = grpc.loadPackageDefinition(packageDefinition).post;
var client = new def['PostService']('localhost:50051', grpc.credentials.createInsecure());
function getPosts() {
    return __awaiter(this, void 0, void 0, function () {
        var call, helper, posts, _i, posts_1, p;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    call = client.getPosts({});
                    helper = function () {
                        return new Promise(function (res, rej) {
                            var data = [];
                            call.on('data', function (post) {
                                data.push(post);
                            });
                            call.on('end', function () { return res(data); });
                            call.on('error', function (err) { return rej(err); });
                        });
                    };
                    return [4 /*yield*/, helper()];
                case 1:
                    posts = _a.sent();
                    for (_i = 0, posts_1 = posts; _i < posts_1.length; _i++) {
                        p = posts_1[_i];
                        console.log(p);
                    }
                    return [2 /*return*/];
            }
        });
    });
}
function createPost() {
    return __awaiter(this, void 0, void 0, function () {
        var post;
        return __generator(this, function (_a) {
            post = {
                title: 'callFromJs',
                content: 'it is an example to invoke create post from js'
            };
            return [2 /*return*/, new Promise(function (resolve, reject) {
                    client.createPost(post, function (err, response) {
                        if (err)
                            return reject(err);
                        console.log(response);
                        resolve(response);
                    });
                })];
        });
    });
}
function updatePost(id) {
    return __awaiter(this, void 0, void 0, function () {
        var post;
        return __generator(this, function (_a) {
            post = {
                id: id,
                title: 'callFromJs',
                content: 'it is an example to invoke create post from js!!!!'
            };
            return [2 /*return*/, new Promise(function (resolve, reject) {
                    client.updatePost(post, function (err, response) {
                        if (err)
                            return reject(err);
                        console.log(response);
                        resolve(response);
                    });
                })];
        });
    });
}
function deletePost(id) {
    return __awaiter(this, void 0, void 0, function () {
        var post;
        return __generator(this, function (_a) {
            post = { id: id };
            return [2 /*return*/, new Promise(function (resolve, reject) {
                    client.deletePost(post, function (err, response) {
                        if (err)
                            return reject(err);
                        console.log(response);
                        resolve(response);
                    });
                })];
        });
    });
}
function main() {
    return __awaiter(this, void 0, void 0, function () {
        var res;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, getPosts()];
                case 1:
                    _a.sent();
                    return [4 /*yield*/, createPost()];
                case 2:
                    res = _a.sent();
                    return [4 /*yield*/, updatePost(res['lastInsertId'])];
                case 3:
                    _a.sent();
                    return [4 /*yield*/, deletePost(res['lastInsertId'])];
                case 4:
                    _a.sent();
                    return [2 /*return*/];
            }
        });
    });
}
main();
