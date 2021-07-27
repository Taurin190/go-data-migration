@startuml
start
:接続確認;
:Srcスキーマ確認;
:Destスキーマ確認;
:SrcとDestのスキーマを比較;
if (同じスキーマが存在する) then (No)
  :ユーザに同じスキーマを作成するか確認する;
  :Destにスキーマを作成する;
endif
:Destのデータを確認;
if (既にデータが存在する) then (Yes)
  :差分更新か全更新か尋ねる;
endif
:Srcのデータ量を確認;
:実行計画を立てる;
repeat :データ移行;
  :一部データ移行を実施;
  :進捗状況を更新;
repeat while (移行完了) is (No) not (Yes)
end
@enduml