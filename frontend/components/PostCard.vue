<template>
  <div>
    <v-row no-gutters>
      <v-col class="mt-2">
        <v-card hover class="px-6" height="200px" @click.native="openPost">
          <v-card-title>{{ post.title }}</v-card-title>
          <v-divider class="mx-3"></v-divider>
          <v-row>
            <v-col cols="8" class="py-0">
              <v-card-subtitle>{{ shortenedPostText }}</v-card-subtitle>
              <v-chip-group class="ml-2" multiple>
                <v-chip v-for="tag in post.tags" :key="tag" dense small pill>
                  {{ tag }}
                </v-chip>
              </v-chip-group>
            </v-col>
            <v-spacer></v-spacer>
            <v-col>
              <author-info></author-info>
            </v-col>
          </v-row>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import AuthorInfo from '@/components/AuthorInfo.vue';

export default {
  components: {
    AuthorInfo,
  },
  props: {
    post: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      votes: 0,
      answers: 0,
      maxTextChars: 100,
    };
  },
  computed: {
    shortenedPostText() {
      const desc = this.post.description;
      return desc.length > this.maxTextChars
        ? `${desc.substring(0, this.maxTextChars)}...`
        : desc;
    },
  },
  methods: {
    openPosts() {
      this.$router.push('/posts/' + this.post.id);
    },
  },
};
</script>

<style lang="scss" scoped>
.v-card__text {
  font-size: 35px;
}
</style>
