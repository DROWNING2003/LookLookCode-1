/**
 * 根据文件类型和名称获取对应的图标类名
 * @param type 文件类型（'dir' 或 'file'）
 * @param name 文件名
 * @returns Font Awesome 图标类名
 */
export const getFileIcon = (type: string, name: string) => {
    if (type === "dir") return "fas fa-folder text-blue-300";
  
    // 根据文件扩展名返回不同图标
    const ext = name.split(".").pop()?.toLowerCase();
    switch (ext) {
      case "md":
        return "fab fa-markdown text-blue-400";
      case "json":
        return "fas fa-file-code text-yellow-400";
      case "js":
      case "ts":
      case "jsx":
      case "tsx":
        return "fab fa-js-square text-yellow-300";
      case "html":
      case "htm":
        return "fab fa-html5 text-orange-500";
      case "css":
      case "scss":
      case "less":
        return "fab fa-css3-alt text-blue-500";
      case "py":
        return "fab fa-python text-blue-400";
      case "java":
        return "fab fa-java text-red-500";
      case "php":
        return "fab fa-php text-purple-500";
      case "go":
        return "fab fa-golang text-blue-500";
      case "rb":
        return "fas fa-gem text-red-500";
      case "sh":
        return "fas fa-terminal text-green-500";
      case "png":
      case "jpg":
      case "jpeg":
      case "gif":
      case "svg":
        return "fas fa-image text-purple-400";
      case "zip":
      case "tar":
      case "gz":
      case "rar":
        return "fas fa-file-archive text-yellow-500";
      case "pdf":
        return "fas fa-file-pdf text-red-500";
      case "doc":
      case "docx":
        return "fas fa-file-word text-blue-500";
      case "xls":
      case "xlsx":
        return "fas fa-file-excel text-green-500";
      case "ppt":
      case "pptx":
        return "fas fa-file-powerpoint text-orange-500";
      case "mp3":
      case "wav":
      case "ogg":
        return "fas fa-file-audio text-purple-400";
      case "mp4":
      case "mov":
      case "avi":
        return "fas fa-file-video text-red-400";
      default:
        return "fas fa-file text-gray-400";
    }
  };