package main

var configCn = `
# 配置将在应用第一次启动时写入到数据库的数据表 goadmin_site 中。后续启动将优先从 goadmin_site
# 中进行加载配置，如果希望修改，可以点击网站右上角配置中心入口进入修改。

# 数据库设置，支持配置多个数据库，目前支持的数据库驱动为：sqlite/mssql/mysql/postgresql
# 默认数据库连接名为default，框架中可以通过自定义的数据库连接名获取到该连接对象。
# 在数据表模型中也可以通过指定对应的连接名来获取对应数据。
database:
  default:
{{- $database := (index .Databases "default")}}
    driver: {{$database.Driver}}
{{- if ne $database.Driver "sqlite"}}
    host: {{$database.Host}}
    port: {{$database.Port}}
    user: {{$database.User}}
    pwd: {{$database.Pwd}}
    name: {{$database.Name}}    
    max_idle_con: {{$database.MaxIdleConns}}
    max_open_con: {{$database.MaxOpenConns}}
{{- else}}
    file: {{$database.File}}
{{end}}
    # params为驱动需要的额外的传参
    # params:
    #   character: utf8mb4

    # 如果设置了DSN，那么以上配置除了Driver都将失效而以配置的DSN为准
    # dsn: ""

# 本应用的唯一ID
app_id: {{.AppID}}

# 定义的网站域名，用于cookie认证进行域名限制
# domain: {{.Domain}}
# 网站语言
language: {{.Language}}
# 全局路由前缀
prefix: {{.UrlPrefix}}
# UI主题
theme: {{.Theme}}
# 文件存储设置，设置上传文件的存储路径以及路由前缀
store:
  path: {{.Store.Path}}
  prefix: {{.Store.Prefix}}

# 网站标题
title: {{.Title}}
# 网站LOGO文字，将显示在登录页面以及侧边栏上方，可以为自定义HTML
logo: {{.Logo}}
# 网站LOGO缩小文字，将显示缩小的侧边栏上方，可以为自定义HTML
mini_logo: {{.MiniLogo}}
# 首页路由
index: {{.IndexUrl}}
# 登录路由
login_url: {{.LoginUrl}}

# 是否为调试模式
debug: {{.Debug}}
# 开发环境：本地 EnvLocal / 测试 EnvTest / 生产 EnvProd
env: {{.Env}}

# info日志本地存储路径
info_log: {{.InfoLogPath}}
# error日志本地存储路径
error_log: {{.ErrorLogPath}}
# access日志本地存储路径
access_log: {{.AccessLogPath}}

# 是否关闭资源访问日志
# access_assets_log_off: {{.AccessAssetsLogOff}}
# 是否关闭sql日志
# sql_log: {{.SqlLog}}
# 是否关闭access日志
# access_log_off: {{.AccessLogOff}}
# 是否关闭info日志
# info_log_off: {{.InfoLogOff}}
# 是否关闭error日志
# error_log_off: {{.ErrorLogOff}}

# 颜色主题，当框架主题为adminlte时生效
# color_scheme: {{.ColorScheme}}

# session的时长，单位为秒，默认为两小时。连续不登两小时后需要重新登录。
session_life_time: {{.SessionLifeTime}}

# 资源路由，当使用CDN时，此配置生效
# asset_url: {{.AssetUrl}}

# 文件上传引擎
file_upload_engine:
  name: {{.FileUploadEngine.Name}}

# 自定义头部HTML
# custom_head_html: {{.CustomHeadHtml}}
# 自定义底部HTML
# custom_foot_html: {{.CustomFootHtml}}
# 自定义底部信息
# footer_info: {{.FooterInfo}}

# 登录页标题
# login_title: {{.LoginTitle}}
# 登录页Logo
# login_logo: {{.LoginLogo}}

# 自定义的用户表
# auth_user_table: {{.AuthUserTable}}

# 是否不限制多IP登录，如果需要多浏览器登录，请设置为true
# no_limit_login_ip: {{.NoLimitLoginIP}}

# 是否关闭网站
# site_off: {{.SiteOff}}

# 是否隐藏配置中心入口
# hide_config_center_entrance: {{.HideConfigCenterEntrance}}
# 是否禁止配置修改
# prohibit_config_modification: {{.ProhibitConfigModification}}
# 是否隐藏应用中心入口
# hide_app_info_entrance: {{.HideAppInfoEntrance}}
# 是否隐藏工具入口
# hide_tool_entrance: {{.HideToolEntrance}}
# 是否隐藏插件中心入口
# hide_plugin_entrance: {{.HidePluginEntrance}}

# 自定义404页面HTML
# custom_404_html: {{.Custom404HTML}}
# 自定义403页面HTML
# custom_403_html: {{.Custom403HTML}}
# 自定义500页面HTML
# custom_500_html: {{.Custom500HTML}}

# 是否开放admin api
# open_admin_api: {{.OpenAdminApi}}
# 是否隐藏用户中心入口
# hide_visitor_user_center_entrance: {{.HideVisitorUserCenterEntrance}}

# 排除的需要加载的主题组件
# exclude_theme_components:
# - ""

# 引导文件的本地路径
bootstrap_file_path: {{.BootstrapFilePath}}
# go.mod文件的本地路径
go_mod_file_path: {{.GoModFilePath}}

# 是否允许删除操作日志
allow_del_operation_log: {{.AllowDelOperationLog}}
# 是否关闭操作日志
operation_log_off: {{.OperationLogOff}}

# 资源文件的本地路径
# 当选择资源文件分离的主题模式时候需要设置此配置项。
{{if eq .AssetRootPath ""}}# {{end}}asset_root_path: {{.AssetRootPath}}

# URL格式
# url_format:
#   info: /info/:__prefix
#   detail: /info/:__prefix/detail
#   create: /new/:__prefix
#   delete: /delete/:__prefix
#   export: /export/:__prefix
#   edit: /edit/:__prefix
#   show_edit: /info/:__prefix/edit
#   show_create: /info/:__prefix/new
#   update: /update/:__prefix`

var configEn = `
# The configuration will be written to the database table goadmin_site when the application
# is first started. Subsequent startup will give priority to loading configuration from
# goadmin_site. If you want to modify, you can click the configuration center entry button in the
# upper right corner of the website to finish the modification.

# Database settings, support the configuration of multiple databases, currently supported
# database drivers are: sqlite/mssql/mysql/postgresql.
# The default database connection name is default, and the connection object can be obtained
# through a custom database connection name in the framework.
# In the data table model, you can also obtain the corresponding data by specifying the
# corresponding connection name.
database:
  default:
{{- $database := (index .Databases "default")}}
    driver: {{$database.Driver}}
{{- if ne $database.Driver "sqlite"}}
    host: {{$database.Host}}
    port: {{$database.Port}}
    user: {{$database.User}}
    pwd: {{$database.Pwd}}
    name: {{$database.Name}}    
    max_idle_con: {{$database.MaxIdleCon}}
    max_open_con: {{$database.MaxOpenCon}}
{{- else}}
    file: {{$database.File}}
{{end}}
    # params are additional parameters required by the driver.
    # params:
    #   character: utf8mb4

    # If DSN is set, all the above configurations except Driver will be invalid and the
    # configured DSN shall prevail
    # dsn: ""

# Unique ID of this application.
app_id: {{.AppID}}

# Defined website domain name, used for cookie authentication for domain name restriction.
# domain: {{.Domain}}
# Language of the website
language: {{.Language}}
# Global url prefix
prefix: {{.UrlPrefix}}
# UI theme
theme: {{.Theme}}
# File storage settings, set the storage path and routing prefix of uploaded files.
store:
  path: {{.Store.Path}}
  prefix: {{.Store.Prefix}}

# Title of the website
title: {{.Title}}
# Website LOGO text, will be displayed above the sidebar, can be custom HTML.
logo: {{.Logo}}
# The website LOGO shrinks the text, and it will display the reduced sidebar above the
# sidebar, which can be customized HTML.
mini_logo: {{.MiniLogo}}
# Index page url.
index: {{.IndexUrl}}
# Login page url.
login_url: {{.LoginUrl}}

# Whether it is debug mode.
debug: {{.Debug}}
# Develop environment: EnvLocal/EnvTest/EnvProd
env: {{.Env}}

# Local storage path of info log.
info_log: {{.InfoLogPath}}
# Local storage path of error log.
error_log: {{.ErrorLogPath}}
# Local storage path of access log.
access_log: {{.AccessLogPath}}

# Whether to close the assets access log.
# access_assets_log_off: {{.AccessAssetsLogOff}}
# Whether to close the sql log.
# sql_log: {{.SqlLog}}
# Whether to close the access log.
# access_log_off: {{.AccessLogOff}}
# Whether to close the info log.
# info_log_off: {{.InfoLogOff}}
# Whether to close the error log.
# error_log_off: {{.ErrorLogOff}}

# Color theme, effective when the theme is adminlte.
# color_scheme: {{.ColorScheme}}

# The duration of the session, in seconds, the default is two hours. You need to log
# in again after not logging in for two hours.
session_life_time: {{.SessionLifeTime}}

# Assets url when using CDN.
# asset_url: {{.AssetUrl}}

# File upload engine.
file_upload_engine:
  name: {{.FileUploadEngine.Name}}

# Custom head HTML.
# custom_head_html: {{.CustomHeadHtml}}
# Custom foot HTML.
# custom_foot_html: {{.CustomFootHtml}}
# Custom footer info HTML.
# footer_info: {{.FooterInfo}}

# Title of the login page.
# login_title: {{.LoginTitle}}
# Logo of the login page.
# login_logo: {{.LoginLogo}}

# Custom auth user table.
# auth_user_table: {{.AuthUserTable}}

# Whether to not restrict multi-IP login, if you need multi-browser login, please set to true.
# no_limit_login_ip: {{.NoLimitLoginIP}}

# Whether to close the website.
# site_off: {{.SiteOff}}

# Whether to hide the configuration center entrance.
# hide_config_center_entrance: {{.HideConfigCenterEntrance}}
# Whether to prohibit configuration modification.
# prohibit_config_modification: {{.ProhibitConfigModification}}
# Whether to hide the application center entrance.
# hide_app_info_entrance: {{.HideAppInfoEntrance}}
# Whether to hide the tool entrance.
# hide_tool_entrance: {{.HideToolEntrance}}
# Whether to hide the entrance to the plug-in center.
# hide_plugin_entrance: {{.HidePluginEntrance}}

# Custom 404 page HTML.
# custom_404_html: {{.Custom404HTML}}
# Custom 403 page HTML.
# custom_403_html: {{.Custom403HTML}}
# Customize 500 pages HTML.
# custom_500_html: {{.Custom500HTML}}

# Whether to open admin api.
# open_admin_api: {{.OpenAdminApi}}
# Whether to hide the user center entrance.
# hide_visitor_user_center_entrance: {{.HideVisitorUserCenterEntrance}}

# Excluded theme components that need to be loaded.
# exclude_theme_components:
# - ""

# The local path of the boot file
bootstrap_file_path: {{.BootstrapFilePath}}
# The local path of the go.mod file
go_mod_file_path: {{.GoModFilePath}}

# Whether to allow deletion of operation log.
allow_del_operation_log: {{.AllowDelOperationLog}}
# Whether to close the operation log.
operation_log_off: {{.OperationLogOff}}

# The local path of the resource file.
# This configuration item needs to be set when the theme mode with resource file separation is selected.
{{if eq .AssetRootPath ""}}# {{end}}asset_root_path: {{.AssetRootPath}}

# URL formats.
# url_format:
#   info: /info/:__prefix
#   detail: /info/:__prefix/detail
#   create: /new/:__prefix
#   delete: /delete/:__prefix
#   export: /export/:__prefix
#   edit: /edit/:__prefix
#   show_edit: /info/:__prefix/edit
#   show_create: /info/:__prefix/new
#   update: /update/:__prefix`
