import { defineStore } from "pinia";

// 定义主题类型（TypeScript）
type ThemeMode = "default" | "customer";
/**
 * 定义主题样式属性的接口，包含必选属性及允许添加额外自定义属性
 */
export interface Theme {
  // 主按钮、链接、高亮元素的颜色
  "--primary-color": string;
  // 次级按钮或辅助色
  "--secondary-color": string;
  // 强调色，用于警告、错误等场景
  "--accent-color": string;
  // 背景颜色
  "--background": string;
  //透明度
  "--background-alpha": string;
  //模糊度
  "--blur": string;
  //背景图片
  "--background-image": string;

  // 主要文本颜色
  "--text-primary": string;
  // 次要文本颜色
  "--text-secondary": string;
  // 默认边框颜色
  "--border-color": string;
   // 默认边框宽度
   "--border-width": string;
  // 鼠标悬停时的边框颜色
  "--border-color-hover": string;
  // 默认的阴影样式
  "--box-shadow": string;
  // 默认的圆角半径
  "--border-radius": string;
  // 鼠标悬停时的阴影样式
  "--box-shadow-hover": string;
  // 鼠标悬停时的圆角半径
  "--border-radius-hover": string;
  // 过渡动画的速度
  "--transition-speed": string;
  // 过渡动画的曲线类型
  "--transition-easing": string;
  // 允许添加任意多的自定义属性，键为字符串，值也为字符串
  [key: string]: string;
}

export const useThemeStore = defineStore("theme", {
  persist: true,
  state: () => ({
    mode: "default" as ThemeMode, //当前主题
    theme: {
      default:<Theme>{
        "--primary-color": "#8365f1", //主按钮、链接、高亮元素
        "--secondary-color": "#2ecc71", //次级按钮/辅助色
        "--accent-color": "#e74c3c", //强调色（警告、错误）

        "--background": "#111722", //背景色
        "--background-alpha": "90%", //透明度
        "--blur": "0px", //模糊度
        "--background-image":"url(@/../default.jpg)", //背景图

        "--text-primary": "#ffff", //文本颜色
        "--front-size": "1.2rem", //文本大小
        "--text-secondary": "#d6d6d6", //次文本颜色
        //边框
        "--border-color": "#e0e0e0", //默认边框
        "--border-width": "0px", //默认边框宽度
        "--border-color-hover": "#bdc3c7", //鼠标悬停时的边框
        "--box-shadow": "0 2px 4px rgba(0,0,0,0.1)", //阴影
        "--border-radius": "20px", //默认圆角
        "--box-shadow-hover": "0 4px 8px rgba(0,0,0,0.1)", //鼠标悬停时的边框
        "--border-radius-hover": "6px", //鼠标悬停时的边框
        //过渡效果
        "--transition-speed": "0.3s", //过渡速度
        "--transition-easing": "ease-in-out" /* 动画曲线 */,
      },
      customer: <Theme>{
        "--primary-color": "#3498db", //主按钮、链接、高亮元素
        "--secondary-color": "#2ecc71", //次级按钮/辅助色
        "--accent-color": "#e74c3c", //强调色（警告、错误）
    
        "--background": "#111722", //背景色
        "--background-alpha": "90%", //透明度
        "--blur": "0px", //模糊度
        "--background-image":"url(@/../public/default.jpg)", //背景图

        "--text-primary": "#333", //文本颜色
        "--text-secondary": "#333", //文本颜色
        //边框
        "--border-color": "#e0e0e0", //默认边框
        "--border-width": "2px", //默认边框宽度
        "--border-color-hover": "#bdc3c7", //鼠标悬停时的边框
        "--box-shadow": "0 2px 4px rgba(0,0,0,0.1)", //阴影
        "--border-radius": "4px", //默认圆角
        "--box-shadow-hover": "0 4px 8px rgba(0,0,0,0.1)", //鼠标悬停时的边框
        "--border-radius-hover": "6px", //鼠标悬停时的边框
        //过渡效果
        "--transition-speed": "0.3s", //过渡速度
        "--transition-easing": "ease-in-out" /* 动画曲线 */,
      },
    },
  }),
  actions: {
    // 切换主题模式
    toggleTheme(mode: ThemeMode) {
      this.mode = mode;
      this.applyTheme();
    },
    // 将 CSS 变量应用到页面根元素
    applyTheme() {
      const themeColors = this.theme[this.mode];
      Object.entries(themeColors).forEach(([key, value]) => {
        document.documentElement.style.setProperty(key, value);
      });
    },
    
    //自定义主题
    setCustomTheme(theme: Theme) {
      this.mode = "customer";
      this.theme.customer = theme;
      this.applyTheme();
    },
  },
  getters: {},
});
