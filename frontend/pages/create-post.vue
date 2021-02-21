<template>
  <v-card class="pa-8">
    <v-card-title class="mb-4">Create Post</v-card-title>
    <v-text-field
      v-model="post.title"
      label="Post title"
      :rules="inputRules"
      outlined
    ></v-text-field>
    <v-textarea
      v-model="post.content"
      label="Descirbe your issue"
      :rules="inputRules"
      outlined
    ></v-textarea>
    <v-card-actions>
      <v-btn class="success" @click.prevent="submitPost">Submit</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      post: {
        title: '',
        content: '',
      },
      inputRules: [
        (v) => !!v || 'Title is required',
        (v) => (v && v.length >= 6) || 'Title must be more than 6 characters',
      ],
    };
  },
  methods: {
    submitPost() {
      console.log(this.post);
    },
  },
};
</script>
