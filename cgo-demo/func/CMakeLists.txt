aux_source_directory(. DIR_LIB_SRCS)

# 生成静态链接库
add_library (add ${DIR_LIB_SRCS})


# add_library (add SHARED ${DIR_LIB_SRCS})
