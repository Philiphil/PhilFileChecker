//com | "'` safe
/((\/\/)(?=(?:[^"'`]*"[^"'`]*"[^"'`]*|[^"'`])*$)(.*))/g
________________
#com
/((#)(.*))/g
_______________________________________
/((\/\/)(?=(?:[^"'`]*"[^"'`]*"[^"'`]*|[^"'`])*$)(.*))/g
/(((\/\/|#)(.*)?(\r|\n))|(([\/\*])(.*)?(\*\/))|((\<\!\-\-)(.*)?(\-\-\>)))/mg
/(((\/\/|#)(?=(?:[^("|'|`)]*("|'|`)[^("|'|`)]*("|'|`)[^("|'|`)]*|[^("|'|`)])*$)(.*)?(\r|\n))|(([\/\*])(.*)?(\*\/))|((\<\!\-\-)(.*)?(\-\-\>)))/mg
/(((\/\/|#)(?=(?:[^("|'|`)]*("|'|`)[^("|'|`)]*("|'|`)[^("|'|`)]*|[^("|'`)])*$)(.*)?(\r|\n))|(([\/\*])(?=(?:[^("|'|`)]*("|'|`)[^("|'|`)]*("|'|`)[^("|'|`)]*|[^("|'|`)])*$)(.*)?(\*\/))|((\<\!\-\-)(?=(?:[^("|'|`)]*("|'|`)[^("|'|`)]*("|'|`)[^("|'|`)]*|[^("|'|`)])*$)(.*)?(\-\-\>)))/mg
/((\/\/)(?=(?:[^("|'|`)]*("|'|`)[^("|'|`)]*("|'|`)[^("|'|`)]*|[^("|'|`)])*$))/gm



/(((\/\*)(?=(?:[^("|'|`)]*("|'|`)[^("|'|`)]*("|'|`)[^("|'|`)]*|[^("|'|`)])*$)(.*?)(\*\/)))/m
