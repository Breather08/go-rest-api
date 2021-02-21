export const state = () => ({
  posts: [],
});

export const mutations = {
  addPost(state, posts) {
    state.posts = posts;
  },
};

export const actions = {};

export const getters = {
  posts: (state) => state.posts,
};
