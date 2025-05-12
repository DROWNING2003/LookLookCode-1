import { useRouter } from "vue-router";

const router = useRouter();
/**
 * 处理文件导航
 * @param file 选中的文件或目录
 */
export const navigateToFile = (file:any,owner: any,repo: any,branch: any,path: any) => {
    console.log("文件详情:", file);
    // 如果是文件，获取文件内容
    if (file.type === "file") {
      
      // getRepositoryContents(owner, repo, file.path as string, currentBranch.value)
      //   .then((response) => {
      //     const content = response.data;
      //     console.log("文件内容:", content);
      //   })
      //   .catch((err) => {
      //     console.error("获取文件内容失败:", err);
      //     error.value = "获取文件内容失败，请稍后重试";
      //   });
    }
    if (file.type === "dir") {
      // 如果是目录，更新当前路径并重新获取内容
      router.push({
        name: "RepoTree",
        params: { owner, repo,branch, path },
      });
      // currentPath.value = file.path as string;
      // console.log(currentPath.value);
      
      // getRepositoryContents(owner, repo, currentPath.value, currentBranch.value)
      //   .then((response) => {
      //     files.value = Array.isArray(response.data)
      //       ? response.data
      //       : [response.data];
      //   })
      //   .catch((err) => {
      //     console.error("Error fetching directory contents:", err);
      //     error.value = "获取目录内容失败，请稍后重试";
      //   });
    } else {
      // 获取文件内容
    }
  };