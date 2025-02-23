package utils

var TempMain = `#include <stdio.h>

#include <FreeRTOS.h>
#include <task.h>

#include <bl_gpio.h>
#include <blog.h>


int main(void){

  // add your code here


  while(1){

  }
    return 0;
}
`

var TempMakeFile = `#
# This is a project Makefile. It is assumed the directory this Makefile resides in is a
# project subdirectory.
#

PROJECT_NAME := $(PROJECTNAME)
PROJECT_PATH := $(abspath .)
PROJECT_BOARD := evb
export PROJECT_PATH PROJECT_BOARD
#CONFIG_TOOLPREFIX :=

-include ./proj_config.mk

ifeq ($(origin BL60X_SDK_PATH), undefined)
BL60X_SDK_PATH ?= $(SDK)
$(info ****** Please SET BL60X_SDK_PATH ******)
$(info ****** Trying SDK PATH [$(BL60X_SDK_PATH)])
endif

COMPONENTS_BLSYS   := bltime blfdt blmtd bloop loopset looprt
COMPONENTS_VFS     := romfs

INCLUDE_COMPONENTS += freertos_riscv_ram
INCLUDE_COMPONENTS += bl602 bl602_std
INCLUDE_COMPONENTS += hosal mbedtls_lts lwip cli vfs yloop utils blog blog_testc newlibc
INCLUDE_COMPONENTS += $(COMPONENTS_NETWORK)
INCLUDE_COMPONENTS += $(COMPONENTS_BLSYS)
INCLUDE_COMPONENTS += $(COMPONENTS_VFS)
INCLUDE_COMPONENTS += $(PROJECT_NAME)

include $(BL60X_SDK_PATH)/make_scripts_riscv/project.mk

`

var TempProjectConfig = `####
CONFIG_SYS_VFS_ENABLE:=1
CONFIG_SYS_VFS_UART_ENABLE:=1
CONFIG_SYS_AOS_CLI_ENABLE:=1
CONFIG_SYS_AOS_LOOP_ENABLE:=1
CONFIG_SYS_BLOG_ENABLE:=1
CONFIG_SYS_DMA_ENABLE:=1
CONFIG_SYS_USER_VFS_ROMFS_ENABLE:=0
CONFIG_SYS_APP_TASK_STACK_SIZE:=4096
CONFIG_SYS_APP_TASK_PRIORITY:=15

CONFIG_BL602_USE_ROM_DRIVER:=1
CONFIG_LINK_ROM=1
CONFIG_FREERTOS_TICKLESS_MODE:=0
CONFIG_WIFI:=0

LOG_ENABLED_COMPONENTS:= blog_testc hosal $(PROJECTNAME)
`

var TempBouffalo = `#
# "main" pseudo-component makefile.
#
# (Uses default behaviour of compiling all source files in directory, adding 'include' to include path.)

include $(BL60X_SDK_PATH)/components/network/ble/ble_common.mk

ifeq ($(CONFIG_ENABLE_PSM_RAM),1)
CPPFLAGS += -DCONF_USER_ENABLE_PSRAM
endif

ifeq ($(CONFIG_ENABLE_CAMERA),1)
CPPFLAGS += -DCONF_USER_ENABLE_CAMERA
endif

ifeq ($(CONFIG_ENABLE_BLSYNC),1)
CPPFLAGS += -DCONF_USER_ENABLE_BLSYNC
endif

ifeq ($(CONFIG_ENABLE_VFS_SPI),1)
CPPFLAGS += -DCONF_USER_ENABLE_VFS_SPI
endif

ifeq ($(CONFIG_ENABLE_VFS_ROMFS),1)
CPPFLAGS += -DCONF_USER_ENABLE_VFS_ROMFS
endif
`
