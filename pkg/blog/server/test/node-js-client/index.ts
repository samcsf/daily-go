import * as path from 'path';
import * as grpc from 'grpc';
import * as protoLoader from '@grpc/proto-loader';

const PROTO_PATH = path.join(__dirname, '../../grpc/post/post/post.proto');

const packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
  }
);

const def = grpc.loadPackageDefinition(packageDefinition).post;
const client = new def['PostService']('localhost:50051', grpc.credentials.createInsecure());

async function getPosts() {
  let call = client.getPosts({});
  let helper = (): Promise<Array<Object>> => {
    return new Promise((res, rej) => {
      let data = [];
      call.on('data', post => {
        data.push(post);
      });
      call.on('end', () => res(data))
      call.on('error', err => rej(err))
    })
  };
  let posts = await helper();
  for (let p of posts) {
    console.log(p);
  }
}

async function createPost() {
  let post = {
    title: 'callFromJs',
    content: 'it is an example to invoke create post from js',
  };
  return new Promise((resolve, reject)=>{
    client.createPost(post, (err, response) => {
      if (err) return reject(err);
      console.log(response);
      resolve(response);
    });
  });
}

async function updatePost(id) {
  let post = {
    id,
    title: 'callFromJs',
    content: 'it is an example to invoke create post from js!!!!',
  };
  return new Promise((resolve, reject)=>{
    client.updatePost(post, (err, response) => {
      if (err) return reject(err);
      console.log(response);
      resolve(response);
    });
  });
}

async function deletePost(id) {
  let post = { id };
  return new Promise((resolve, reject)=>{
    client.deletePost(post, (err, response) => {
      if (err) return reject(err);
      console.log(response);
      resolve(response);
    });
  });
}

async function main() {
  await getPosts();
  let res = await createPost();
  await updatePost(res['lastInsertId']);
  await deletePost(res['lastInsertId']);
}

main();
