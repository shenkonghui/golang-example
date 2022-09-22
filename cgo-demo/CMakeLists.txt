cmake_minimum_required(VERSION 2.6.4)

set( project_name "HelloWorld")


project (${project_name})

aux_source_directory(. DIR_SRCS)

add_executable(${project_name} ${DIR_SRCS})
