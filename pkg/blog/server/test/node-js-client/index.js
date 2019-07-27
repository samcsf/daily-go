const path = require('path');
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = path.join(__dirname, '../../grpc/post/post/post.proto');
console.log(path.resolve(PROTO_PATH));

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
const post = grpc.loadPackageDefinition(packageDefinition).post;
const client = new post.PostService('localhost:50051', grpc.credentials.createInsecure());

async function getPosts() {
  let call = client.getPosts({});
  let helper = () => {
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

async function updatePost() {
  let post = {
    id: '14',
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

async function deletePost() {
  let post = {
    id: 13,
  };
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
  await createPost();
  await updatePost();
  await deletePost();
}

main();
