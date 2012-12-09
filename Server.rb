require 'webrick'
include WEBrick

class Simple < WEBrick::HTTPServlet::AbstractServlet
  
  def do_GET(request, response)
    response.status = 200
    response['Content-Type'] = "text/plain"
    response.body = "WeeebRick"
  end
  def do_POST(request, response)
    body = String.new
    request.body do |chunk|
      body << chunk
    end

    puts body
       
    response.status = 200
    response['Content-Type'] = "text/plain"
    response.body = "WeeebRick POST"
  end
end

server = HTTPServer.new(
      :Port            => 8000,
      :DocumentRoot    => './'
    )
server.mount "/", Simple

trap("INT"){ server.shutdown }

server.start

#https://github.com/betten/SoundCloud-Developer-Challenge/blob/master/server.rb