import React from 'react'
import { MdFace, MdFileDownload, MdQuestionAnswer } from 'react-icons/md'
import { ME } from '../constants'
import { getNickname } from '../nickname'
import { Message } from '../reducers/messages'
import { Nicknames } from '../reducers/nicknames'
import Input from './Input'

export interface MessageProps {
  message: Message
}

function MessageEntry (props: MessageProps) {
  const { message } = props

  const handleClick = () => {
    const a = document.createElement('a')
    a.setAttribute('href', message.data!)
    a.setAttribute('download', message.message)
    a.style.visibility = 'hidden'
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
  }

  return (
    <p className='message-text'>
      {message.image && (
        <img src={message.data} width='100%' />
      )}
      {message.data && (
        <button
          className='message-download'
          onClick={handleClick}
        >
          <span>{message.message}</span>
          <MdFileDownload className='icon' />
        </button>
      )}
      {!message.data && message.message}
    </p>
  )
}

export interface ChatProps {
  messages: Message[]
  nicknames: Nicknames
  visible: boolean
  sendFile: (file: File) => void
  sendText: (message: string) => void
}

export default class Chat extends React.PureComponent<ChatProps> {
  chatHistoryRef = React.createRef<HTMLDivElement>()
  inputRef = React.createRef<Input>()

  scrollToBottom = () => {
    const chatHistoryRef = this.chatHistoryRef.current!
    chatHistoryRef.scrollTop = chatHistoryRef.scrollHeight
  }
  componentDidMount () {
    this.focus()
    this.scrollToBottom()
  }
  componentDidUpdate () {
    this.focus()
    this.scrollToBottom()
  }
  focus() {
    if (this.props.visible) {
      this.inputRef.current?.textArea.current?.focus()
    }
  }
  render () {
    const { messages, sendFile, sendText } = this.props
    return (
      <div className='chat'>
        <div className='chat-history' ref={this.chatHistoryRef}>

          {messages.length ? (
            messages.map((message, i) => (
              <div key={i}>
                {message.peerId === ME ? (
                  <div className='chat-item chat-item-me'>
                    <div className='message'>
                      <span className='message-user-name'>
                        {getNickname(this.props.nicknames, message.peerId)}
                      </span>
                      <time className='message-time'>{message.timestamp}</time>
                      <MessageEntry message={message} />
                    </div>
                    <span className='chat-item-img'>
                      <MdFace />
                    </span>
                  </div>
                ) : (
                  <div className='chat-item chat-item-other'>
                    <span className='chat-item-img'>
                      <MdFace />
                    </span>
                    <div className='message'>
                      <span className='message-user-name'>
                        {getNickname(this.props.nicknames, message.peerId)}
                      </span>
                      <time className='message-time'>{message.timestamp}</time>
                      <MessageEntry message={message} />
                    </div>
                  </div>
                )}
              </div>
            ))
          ) : (
            <div className='chat-empty'>
              <span className='chat-empty-icon'>
                <MdQuestionAnswer />
              </span>
              <div className='chat-empty-message'>No Notifications</div>
            </div>
          )}

        </div>

        <Input
          ref={this.inputRef}
          sendText={sendText}
          sendFile={sendFile}
        />
      </div>
    )
  }
}
