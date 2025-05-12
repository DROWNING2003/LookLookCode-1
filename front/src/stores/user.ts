import { defineStore } from 'pinia'

export interface UserState {
  VITE_BACK_URL: string,
  id: number
  login: string
  name: string | null
  avatar_url: string
  bio: string | null
  blog: string | null
  company: string | null
  location: string | null
  email: string | null
  followers: number
  following: number
  created_at: string
  updated_at: string
  public_repos: number
  html_url: string
  accessToken:string
}

export const useUserStore = defineStore('user', {
    persist: true,
  state: (): UserState => ({
    VITE_BACK_URL: import.meta.env.VITE_BACK_URL || "http://localhost:3000",
    id: 0,
    login: '',
    name: null,
    avatar_url: '',
    bio: null,
    blog: null,
    company: null,
    location: null,
    email: null,
    followers: 0,
    following: 0,
    created_at: '',
    updated_at: '',
    public_repos: 0,
    html_url: '',
    accessToken:'',
  }),
  getters: {
    // AccessToken: (state) => state.count * 2
  },
  actions: {
    setUser(userData: Partial<UserState>) {
      Object.assign(this, userData)
    },
    clearUser() {
      this.$reset()
    },
    setAccessToken(token:string){
      this.accessToken=token
    },
     // 设置后端地址
     setBackendUrl(url: string) {
      this.VITE_BACK_URL = url;
    },
  },
})