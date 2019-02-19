<!DOCTYPE html>
<html>
<body>
  <form method="POST">
  <table border="1">
    <tbody>
      <tr>
        <th>yaml</th><td style="vertical-align: top">{{or .yaml "New entry" | nl2brAndNbsp}}</td><td><textarea name="yaml" rows="30" cols="60">{{or .yaml ""}}</textarea></td>
      </tr>
    </tbody>
  </table>
  <input type="submit" value="送信"></input>
  <br><a href="/admin/list">戻る</a>
  </form>
</body>
</html>