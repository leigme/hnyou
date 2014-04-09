{{define "navbar.tpl"}}
<div class="container">
  <div class="navbar-header">
    {{template "logo.tpl"}}
  </div>
  <div class="navbar-collapse collapse">
    <form class="navbar-form navbar-right">
      <div class="form-group">
        <div class="btn-group">
          <button type="button" class="btn btn-info dropdown-toggle" data-toggle="dropdown">
          选择地区
          <span class="caret"></span>
          </button>
          <ul class="dropdown-menu" role="menu">
            <li><a href="#">长沙市</a></li>
            <li><a href="#">株洲市</a></li>
            <li><a href="#">湘潭市</a></li>
            <li><a href="#">衡阳市</a></li>
            <li><a href="#">邵阳市</a></li>
            <li><a href="#">岳阳市</a></li>
            <li><a href="#">常德市</a></li>
            <li><a href="#">张家界市</a></li>
            <li><a href="#">益阳市</a></li>
            <li><a href="#">娄底市</a></li>
            <li><a href="#">郴州市</a></li>
            <li><a href="#">永州市</a></li>
            <li><a href="#">怀化市</a></li>
            <li><a href="#">湘西自治州</a></li>
            <li class="divider"></li>
          </ul>
        </div>
      </div>
      <div class="form-group">
        <input type="password" placeholder="输入景点" class="form-control">
      </div>
      <button type="submit" class="btn btn-primary">查询</button>
    </form>
  </div><!--/.navbar-collapse -->
</div>
{{end}}