require 'webrick'
include WEBrick

class Simple < WEBrick::HTTPServlet::AbstractServlet
  
  def do_GET(request, response)
    response.status = 200
    response['Content-Type'] = "text/plain"
    response.body = "WeeebRick"
  end

  def parse_query(content_type, body)
    boundary = content_type.match(/^multipart\/form-data; boundary=(.+)/)[1]
    boundary = HTTPUtils::dequote(boundary)
    return HTTPUtils::parse_form_data(body, boundary)
  end

  def do_POST(request, response)
    body = String.new
    request.body do |chunk|
      body << chunk
    end

#    puts request.raw_header
#    puts body

    filedata = parse_query(request['content-type'], body)
    puts filedata['path']
    puts "----------"
    puts filedata['upload']
                
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