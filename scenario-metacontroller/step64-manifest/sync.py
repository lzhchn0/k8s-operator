from http.server import BaseHTTPRequestHandler, HTTPServer
import io
import json
import copy
import re



Def new_dep(old):
  mydep = copy.deepcopy(old)
  mydep[‘spec’][‘template’]

  mydep[‘metadata’][‘labels’][‘my-new’] = ‘metacontroller’




  return mydep



class Controller(BaseHTTPRequestHandler):

  def do_POST(self):
    observed = json.loads(self.rfile.read(int(self.headers.get('content-length'))))
    desired = self.sync(observed['parent'], observed['children'])

    self.send_response(200)
    self.send_header('Content-type', 'application/json')
    self.end_headers()
    self.wfile.write(io.BytesIO(json.dumps(desired).encode('utf-8')).getvalue())


HTTPServer(('', 80), Controller).serve_forever()

