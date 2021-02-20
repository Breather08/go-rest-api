<template>
  <v-card flat>
    <v-card-subtitle class="pb-1">{{ action }} {{ time }} </v-card-subtitle>
    <v-row no-gutters>
      <v-col class="pt-2" cols="4">
        <v-img v-if="author.avatar" :src="author.avatar" width="30"></v-img>
        <v-icon v-else large>mdi-account</v-icon>
      </v-col>
      <v-col>
        <nuxt-link
          class="author__info author__nickname"
          :to="`/users/${author.nickname}`"
          @click.native.stop
          >{{ author.nickname }}</nuxt-link
        ><br />
        <v-icon small color="orange">mdi-star-circle</v-icon>
        <span class="author__info">{{ author.reputation }}</span>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
export default {
  props: {
    action: {
      type: String,
      default: 'asked',
    },
    time: {
      type: String,
      default: `Jan 3 '18`,
    },
    author: {
      type: Object,
      default: () => ({
        nickname: 'Breather08',
        avatar: null,
        reputation: 152,
      }),
    },
  },
};
</script>
<style lang="scss" scoped>
.author__info {
  font-size: 12px;
}
</style>
