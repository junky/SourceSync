require 'rest_client'

RestClient.post( 'http://localhost:8000',
  :upload => File.new('C:\Temp\key.txt', 'rb'),
  :path => 'Temp\key.txt'
)