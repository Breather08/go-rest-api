<template>
  <v-app>
    <v-app-bar app color="primary" dark> </v-app-bar>

    <v-main>
      <post-form @updatePost="addPost" />
      <post v-for="post in posts" :key="post.title" :title="post.title" :content="post.content" />
    </v-main>
  </v-app>
</template>

<script>
import axios from 'axios';
import PostForm from './components/PostForm.vue';
import Post from './components/Post.vue';

export default {
  name: 'App',
  components: {
    PostForm,
    Post,
  },
  data() {
    return {
      posts: [],
    };
  },
  methods: {
    addPost(data) {
      axios
        .post('http://localhost:3000/api/create-post', data, {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
          },
        })
        .then(() => {
          this.posts = [...this.posts, data];
        })
        .catch((err) => {
          console.error(err);
        });
    },
  },
  async created() {
    await axios
      .get('http://localhost:3000/api/posts', {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
      })
      .then((response) => {
        this.posts = response.data ? response.data : [];
      });
  },
};
</script>
