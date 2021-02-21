<template>
  <div class="pa-10">
    <v-row no-gutters align="center" class="px-3 mb-2">
      <v-col cols="7">
        <h1>All posts</h1>
        <span>{{ posts.length }} posts total</span>
      </v-col>
      <v-spacer></v-spacer>
      <v-switch
        v-model="ascending"
        flat
        inset
        dense
        :prepend-icon="`mdi-sort-${!ascending ? 'ascending' : 'descending'}`"
        class="mx-5 px-0"
      ></v-switch>
      <v-col>
        <v-btn-toggle v-model="sortBy" background-color="#E8E8E8">
          <v-tooltip
            v-for="sortOption in sortOptions"
            :key="sortOption.title"
            transition="none"
            bottom
          >
            <template #activator="{ on, attrs }">
              <v-btn v-bind="attrs" v-on="on">
                <v-icon>{{ sortOption.icon }}</v-icon>
              </v-btn>
            </template>
            <span>{{ sortOption.title }}</span>
          </v-tooltip>
        </v-btn-toggle>
      </v-col>
    </v-row>
    <v-divider class="mb-5"></v-divider>
    <post-card
      v-for="post in posts"
      :key="post.title"
      :post="post"
      class="mb-5"
    ></post-card>
    <v-hover v-slot="{ hover }">
      <v-btn
        to="/create-post"
        color="success"
        class="ma-5"
        fixed
        bottom
        rounded
        right
        large
      >
        <v-icon absolute>mdi-plus</v-icon>
        <v-expand-x-transition>
          <span v-if="hover">add post</span>
        </v-expand-x-transition>
      </v-btn>
    </v-hover>
  </div>
</template>

<script>
import PostCard from '@/components/PostCard.vue';

export default {
  title: 'Home',
  components: {
    PostCard,
  },
  data() {
    return {
      posts: [
        {
          id: 0,
          title: 'Some title 1',
          description:
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
          tags: ['java', 'script', 'python', 'golang'],
        },
        {
          id: 1,
          title: 'Some title 2',
          description:
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
          tags: ['java', 'script', 'python', 'golang'],
        },
      ],
      ascending: true,
      sortBy: 'none',
      sortOptions: [
        {
          id: 0,
          title: 'Most recent',
          icon: 'mdi-clock-time-five-outline',
          onclick: this.sortByTime(),
        },
        {
          id: 1,
          title: 'Most discussed',
          icon: 'mdi-comment-outline',
          onclick: this.sortByAnswers(),
        },
        {
          id: 2,
          title: 'Most voted',
          icon: 'mdi-swap-vertical',
          onclick: this.sortByVoted(),
        },
      ],
    };
  },
  methods: {
    sortByTime() {},
    sortByAnswers() {},
    sortByVoted() {},
  },
};
</script>
