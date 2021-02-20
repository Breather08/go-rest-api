<template>
  <v-app class="main-app">
    <v-navigation-drawer
      absolute
      color="#E8E8E8"
      expand-on-hover
      clipped
      fixed
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar color="#495464" class="px-10" clipped-left fixed app>
      <v-img
        :src="mainIcon"
        max-width="200"
        position="bottom -20px left 10px"
      ></v-img>
      <v-toolbar-title />
      <v-spacer />
      <v-btn icon>
        <v-icon color="#f4f4f2"> mdi-login</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import mainIcon from '@/static/srak_logo.png';

export default {
  data() {
    return {
      mainIcon,
      items: [
        {
          icon: 'mdi-format-list-bulleted-square',
          title: 'Posts',
          to: '/',
        },
        {
          icon: 'mdi-account',
          title: 'Profile',
          to: '/profile',
        },
        {
          icon: 'mdi-information-outline',
          title: 'About',
          to: '/about',
        },
      ],
    };
  },
};
</script>

<style lang="scss" scoped>
.main-app {
  background: #f4f4f2;
}
</style>
