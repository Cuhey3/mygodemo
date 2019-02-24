<!DOCTYPE html>
<html>
<body>
  <h3>Executable YAML List</h3>
  <a href="/ws" target="_blank">Websocket Monitor</a> | <a href="/admin/edit">New Entry</a>
  <form method="POST">
  <table border="1">
    <thead>
      <tr>
        <th>id</th><th>yaml</th><th>operation</th><th>processes</th>

      </tr>
    </thead>
    <tbody>
      {{range $item := $.yamls}}
      {{$objectIdHex := $item._id | objectIdToHex}}
      <tr>
          <td><a href="/admin/edit/{{$objectIdHex }}">{{$objectIdHex}}</a></td>
          <td>{{$item.yaml | nl2brAndNbsp }}</td>
          <td><button type="submit" name="run" value="{{$objectIdHex}}">run</input></td>
          
          <td>{{range $childId := index $.processes $objectIdHex}}
          {{$childId}}<button type="submit" name="kill" value="{{$objectIdHex}}{{$childId}}">kill</button><br>
          {{end}}</td>
      </tr>
      {{end}}
    </tbody>
  </table>
  </form>
</body>
</html>